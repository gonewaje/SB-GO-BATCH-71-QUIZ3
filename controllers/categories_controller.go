package controllers

import (
	"database/sql"
	"library/repository"
	"library/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoriesController struct {
	DB *sql.DB
}

func (ctl CategoriesController) List(c *gin.Context) {
	cats, err := repository.ListCategories(ctl.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cats})
}

func (ctl CategoriesController) Create(c *gin.Context) {
	var in structs.Category
	if err := c.ShouldBindJSON(&in); err != nil || in.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}
	if err := repository.CreateCategory(ctl.DB, in.Name, "api"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "category created"})
}

func (ctl CategoriesController) Detail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	cat, err := repository.GetCategory(ctl.DB, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cat})
}

func (ctl CategoriesController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	ok, err := repository.DeleteCategory(ctl.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "category deleted"})
}

func (ctl CategoriesController) BooksByCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	books, err := repository.ListBooksByCategory(ctl.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": books})
}
