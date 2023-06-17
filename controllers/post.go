package controllers

import (
	"course-forum/infra/logger"
	"course-forum/models"
	"course-forum/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func getCache(ctx *gin.Context, key string, data interface{}) (err error) {
	var val string

	//get data form redis
	if val, err = repository.RedisGet(ctx, key); err != nil {
		return
	}

	//convert json to struct
	if err = json.Unmarshal([]byte(val), &data); err != nil {
		logger.Errorf("json convert error: %s \n", err.Error())
		go repository.RedisDelete(ctx, key)
	}

	return
}

func setCache(ctx *gin.Context, key string, data interface{}) {
	val, err := json.Marshal(data)

	if err != nil {
		logger.Infof("json convert error: %s \n", err.Error())
		return
	}

	go repository.RedisSet(ctx, key, val)
}

// GetPosts godoc
//
// @Summary Get all posts
// @Description Get all posts
// @Tags posts
// @Accept json
// @Produce json
// @Success 200 {array} models.Post
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /posts [get]
func GetPosts(ctx *gin.Context) {
	var posts []models.Post
	var err error

	//get data form cache
	if err = getCache(ctx, "posts", &posts); err == nil {
		ctx.JSON(http.StatusOK, &posts)
		return
	}

	//get data form database
	if posts, err = repository.GetPosts(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &posts)
	go setCache(ctx, "posts", &posts)
}

// CreatePost godoc
//
// @Summary Create post
// @Description Create a post
// @Tags posts
// @Accept json
// @Produce json
// @Param request body models.CreatePost true "Create post request"
// @Success 201 {object} models.Post
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /posts [post]
func CreatePost(ctx *gin.Context) {
	var input models.CreatePost

	_ = ctx.ShouldBindJSON(&input)

	if err := validate.Struct(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var tags []models.Tag = []models.Tag{}
	var err error
	if len(input.Tags) != 0 {
		tags, err = repository.GetTags(input.Tags)
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{Title: input.Title, Content: input.Content, Score: input.Score, SentimentScore: input.SentimentScore ,CreateBy: input.CreateBy, Tags: tags}

	if err = repository.CreatePost(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, &post)
	go repository.RedisDelete(ctx, "posts")
}

// FindPost godoc
//
// @Summary Find post
// @Description Find a post
// @Tags posts
// @Accept json
// @Produce json
// @Param id path uint true "Post ID"
// @Success 200 {object} models.Post
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /posts/{id} [get]
func FindPost(ctx *gin.Context) {
	id, err := getPostId(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id must be a number"})
		return
	}

	var post *models.Post
	var key string = fmt.Sprintf("posts/%v", id)

	//get data form cache
	if err = getCache(ctx, key, &post); err == nil {
		ctx.JSON(http.StatusOK, &post)
		return
	}

	//get data form database
	if post, err = repository.FindPost(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return
	}

	ctx.JSON(http.StatusOK, &post)
	go setCache(ctx, key, &post)
}

// UpdatePost godoc
//
// @Summary Update post
// @Description Update a post
// @Tags posts
// @Accept json
// @Produce json
// @Param id path uint true "Post ID"
// @Param request body models.UpdatePost true "Update post request"
// @Success 200 {object} models.Post
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /posts/{id} [patch]
func UpdatePost(ctx *gin.Context) {
	id, err := getPostId(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id must be a number"})
		return
	}

	var input models.UpdatePost
	_ = ctx.ShouldBindJSON(&input)

	if err = validate.Struct(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{ID: id, State: *input.State}

	if err = repository.UpdatePost(&post); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return
	}

	ctx.JSON(http.StatusOK, &post)
	go repository.RedisDelete(ctx, "posts")
	go repository.RedisDelete(ctx, fmt.Sprintf("posts/%v", id))
}

// DeletePost godoc
//
// @Summary Delete post
// @Description Delete a post
// @Tags posts
// @Accept json
// @Produce json
// @Param id path uint true "Post ID"
// @Success 204
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /posts/{id} [delete]
func DeletePost(ctx *gin.Context) {
	id, err := getPostId(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id must be a number"})
		return
	}

	post := models.Post{ID: id}

	if err = repository.DeletePost(&post); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
	go repository.RedisDelete(ctx, "posts")
	go repository.RedisDelete(ctx, fmt.Sprintf("posts/%v", id))
}

func getPostId(ctx *gin.Context) (uint, error) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		return 0, err
	}

	return uint(id), nil
}
