package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type StudentInfo struct {
	Name    string
	Phone   string
	Course string
	Message string
}

func homePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func aboutPage(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", nil)
}

// func contactPage(c *gin.Context) {
// 	c.HTML(http.StatusOK, "contact.html", nil)
// }

func contactPage(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.HTML(http.StatusOK, "contact.html", nil)
		return
	}
	text := StudentInfo{
		Name:    c.PostForm("name"),
		Phone:   c.PostForm("phone"),
		Course: c.PostForm("email"),
		Message: c.PostForm("message"),
	}
	c.HTML(http.StatusOK, "contact.html", gin.H{
		"Success": true,
		"Message": text,
	})
}

func main() {
	router := gin.Default()

	// Serve static files
	router.Static("/assets", "./assets")

	// Load templates
	router.LoadHTMLGlob("templates/*.html")

	// Define routes
	router.GET("/", homePage)
	router.GET("/contact", contactPage)
	router.POST("/contact", contactPage)
	router.GET("/about", aboutPage)
	// router.POST("/save-contact-message", contactInfo)

	// Run the server
	router.Run(":8888")
}
