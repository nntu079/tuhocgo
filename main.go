package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/getData", getData)
	router.GET("/getData", getDataPost)

	router.Run()
}

func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hi I am GIN Framework",
	})
}
func getDataPost(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hi I am GIN Framework",
	})
}