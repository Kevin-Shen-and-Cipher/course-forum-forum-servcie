package controllers

import (
	"course-forum/models"
	"course-forum/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func init() {
	validate = validator.New()
}

// GetTags godoc
//
// @Summary Get all tags
// @Description Get all tags
// @Tags tags
// @Accept json
// @Produce json
// @Success 200 {array} models.Tag
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /tags [get]
func GetTags(ctx *gin.Context) {
	var tags []models.Tag
	var err error

	//get data form cache
	if err = getCache(ctx, "tags", &tags); err == nil {
		ctx.JSON(http.StatusOK, &tags)
		return
	}

	//get data form database
	if tags, err = repository.GetTags(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &tags)
	go setCache(ctx, "tags", &tags)
}

// CreateTag godoc
//
// @Summary Create tag
// @Description Create a tag
// @Tags tags
// @Accept json
// @Produce json
// @Param request body models.CreateTag true "Create tag request"
// @Success 201 {object} models.Tag
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /tags [post]
func CreateTag(ctx *gin.Context) {
	var input models.CreateTag

	_ = ctx.ShouldBindJSON(&input)

	if err := validate.Struct(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tag := models.Tag{Name: input.Name}
	err := repository.CreateTag(&tag)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, &tag)
	go repository.RedisDelete(ctx, "tags")
}

// FindTag godoc
//
// @Summary Find tag
// @Description Find a tag
// @Tags tags
// @Accept json
// @Produce json
// @Param id path uint true "Tag ID"
// @Success 200 {object} models.Tag
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /tags/{id} [get]
func FindTag(ctx *gin.Context) {
	id, err := getTagId(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id must be a number"})
		return
	}

	var tag *models.Tag
	var key string = fmt.Sprintf("tags/%v", id)

	//get data form cache
	if err = getCache(ctx, key, &tag); err == nil {
		ctx.JSON(http.StatusOK, &tag)
		return
	}

	//get data form database
	if tag, err = repository.FindTag(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "tag not found"})
		return
	}

	ctx.JSON(http.StatusOK, &tag)
	go setCache(ctx, key, &tag)
}

// UpdateTag godoc
//
// @Summary Update tag
// @Description Update a tag
// @Tags tags
// @Accept json
// @Produce json
// @Param id path uint true "Tag ID"
// @Param request body models.UpdateTag true "Update tag request"
// @Success 200 {object} models.Tag
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /tags/{id} [patch]
func UpdateTag(ctx *gin.Context) {
	id, err := getTagId(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id must be a number"})
		return
	}

	var input models.UpdateTag
	_ = ctx.ShouldBindJSON(&input)

	if err := validate.Struct(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tag := models.Tag{ID: id, Name: input.Name}

	if err = repository.UpdateTag(&tag); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "tag not found"})
		return
	}

	ctx.JSON(http.StatusOK, &tag)
	go repository.RedisDelete(ctx, "tags")
	go repository.RedisDelete(ctx, fmt.Sprintf("tags/%v", id))
}

// DeleteTag godoc
//
// @Summary Delete tag
// @Description Delete a tag
// @Tags tags
// @Accept json
// @Produce json
// @Param id path uint true "Tag ID"
// @Success 204
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /tags/{id} [delete]
func DeleteTag(ctx *gin.Context) {
	id, err := getTagId(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id must be a number"})
		return
	}

	tag := models.Tag{ID: id}

	if err = repository.DeleteTag(&tag); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "tag not found"})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
	go repository.RedisDelete(ctx, "tags")
	go repository.RedisDelete(ctx, fmt.Sprintf("tags/%v", id))
}

func getTagId(ctx *gin.Context) (uint, error) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		return 0, err
	}

	return uint(id), nil
}
