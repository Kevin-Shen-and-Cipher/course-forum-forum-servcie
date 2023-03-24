package controllers

import (
	"course-forum/models"
	"course-forum/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPosts(ctx *gin.Context) {
	posts, err := repository.GetPosts()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, &posts)
	}
}

func CreatePost(ctx *gin.Context) {
	var input models.CreatePost

	if err := ctx.ShouldBindJSON(&input); err != nil {
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

func FindPost(ctx *gin.Context) {
	id, err := getPostId(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id must be a number"})
		return
	}

	post, err := repository.FindPost(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, &post)
	}
}

func UpdatePost(ctx *gin.Context) {
	id, err := getPostId(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id must be a number"})
		return
	}

	var input models.UpdatePost

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "input": input})
		return
	}

	post := models.Post{ID: id, State: input.State}
	err = repository.UpdatePost(&post)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, &post)
	}
}

func DeletePost(ctx *gin.Context) {
	id, err := getPostId(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id must be a number"})
		return
	}

	post := models.Post{ID: id}
	err = repository.DeletePost(&post)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusNoContent, nil)
	}
}

func getPostId(ctx *gin.Context) (uint, error) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		return 0, err
	}

	return uint(id), nil
}
