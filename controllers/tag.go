package controllers

import (
	"course-forum/models"
	"course-forum/repository"
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
	tags, err := repository.GetTags()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, &tags)
	}
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
// @Router /tags [tag]
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
	} else {
		ctx.JSON(http.StatusCreated, &tag)
	}
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

	tag, err := repository.FindTag(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "tag not found"})
	} else {
		ctx.JSON(http.StatusOK, &tag)
	}
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
	err = repository.UpdateTag(&tag)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "tag not found"})
	} else {
		ctx.JSON(http.StatusOK, &tag)
	}
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
	err = repository.DeleteTag(&tag)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "tag not found"})
	} else {
		ctx.JSON(http.StatusNoContent, gin.H{})
	}
}

func getTagId(ctx *gin.Context) (uint, error) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		return 0, err
	}

	return uint(id), nil
}
