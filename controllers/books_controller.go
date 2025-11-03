package controllers

import (
	"database/sql"
	"library/repository"
	"library/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BooksController struct {
	DB *sql.DB
}

func (ctl BooksController) List(c *gin.Context) {
	books, err := repository.ListBooks(ctl.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": books})
}

func (ctl BooksController) Create(c *gin.Context) {
	var in structs.Book
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if in.ReleaseYear < 1980 || in.ReleaseYear > 2024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "release_year must be between 1980 and 2024"})
		return
	}

	if in.TotalPage > 100 {
		in.Thickness = "tebal"
	} else {
		in.Thickness = "tipis"
	}

	if err := repository.CreateBook(ctl.DB, in, "api"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "book created", "thickness": in.Thickness})
}

func (ctl BooksController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	ok, err := repository.DeleteBook(ctl.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "book deleted"})
}

func (bc *BooksController) Detail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book ID"})
		return
	}

	book, err := repository.GetBookByID(bc.DB, id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}
