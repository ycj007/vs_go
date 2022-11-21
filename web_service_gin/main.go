package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "2", Artist: "3", Price: 4},
	{ID: "11", Title: "21", Artist: "31", Price: 41},
	{ID: "12", Title: "22", Artist: "32", Price: 42},
	{ID: "13", Title: "23", Artist: "33", Price: 43},
}

func writeJsonBatch(data *[]album, c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data)
}
func writeSingle(data album, c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data)
}
func writeNotFound(c *gin.Context) {
	c.IndentedJSON(http.StatusNotFound, "not found")
}

func getBatch(c *gin.Context) {
	writeJsonBatch(&albums, c)
}
func postHandler(c *gin.Context) {
	var newData album
	if err := c.BindJSON(&newData); err != nil {
		return
	}
	albums = append(albums, newData)
	writeJsonBatch(&albums, c)
}

func getByPathVarHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		writeNotFound(c)
		return
	}

	for _, v := range albums {
		if v.ID == id {
			writeSingle(v, c)
			return
		}
	}
	writeNotFound(c)

}
func main() {
	engin := gin.Default()
	engin.GET("/albums", getBatch)
	engin.POST("albums", postHandler)
	engin.GET("getById/:id", getByPathVarHandler)
	engin.Run("localhost:8080")
}
