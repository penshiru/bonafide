package main

//InitRouter set routers default and the API endpoints
import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//InitRouter set all the routing paths
func InitRouter(config *Config) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(CORSMiddleware())

	router.POST("/parse", func(c *gin.Context) {

		//read file
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "msg": "Could not upload file"})
			return
		}

		loc, err := SaveUploadedFile(file, file.Filename, "tmp")
		if err != nil {
			c.JSON(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		law := ParseConcurrent(loc)
		c.JSON(http.StatusOK, gin.H{"status": "success", "data": law})

	})

	return router

}

//CORSMiddleware set the CORS headers
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
