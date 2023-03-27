package controllers

import (
	"course-forum/models"
	"course-forum/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
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
	posts, err := repository.GetPosts()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, &posts)
	}
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

	post := models.Post{Title: input.Title, Content: input.Content, Score: input.Score, CreateBy: input.CreateBy}

	err := repository.CreatePost(&post)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &post)
	}
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

	post, err := repository.FindPost(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
	} else {
		ctx.JSON(http.StatusOK, &post)
	}
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

	if err := validate.Struct(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{ID: id, State: *input.State}
	err = repository.UpdatePost(&post)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
	} else {
		ctx.JSON(http.StatusOK, &post)
	}
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
	err = repository.DeletePost(&post)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
	} else {
		ctx.JSON(http.StatusNoContent, gin.H{})
	}
}

func getPostId(ctx *gin.Context) (uint, error) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		return 0, err
	}

	return uint(id), nil
}
