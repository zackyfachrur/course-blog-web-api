package controllers

import (
	"myproject/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Albums)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range models.Albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album tidak ditemukan"})
}

func CreateAlbum(c *gin.Context) {
	var newAlbum models.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	models.Albums = append(models.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func UpdateAlbum(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbum models.Album

	if err := c.BindJSON(&updatedAlbum); err != nil {
		return
	}

	for i, a := range models.Albums {
		if a.ID == id {
			models.Albums[i] = updatedAlbum
			c.IndentedJSON(http.StatusOK, updatedAlbum)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album tidak ditemukan"})
}

func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")
	for i, a := range models.Albums {
		if a.ID == id {
			models.Albums = append(models.Albums[:i], models.Albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Album berhasil dihapus"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album tidak ditemukan"})
}
