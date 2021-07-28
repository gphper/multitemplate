package main

import (
	"embed"
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

//go:embed templates
var templatePath embed.FS

func main() {
	router := gin.Default()
	router.HTMLRender = loadTemplates()
	router.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin.html", gin.H{
			"title": "Welcome!",
		})
	})
	router.GET("/article", func(c *gin.Context) {
		c.HTML(200, "article.html", gin.H{
			"title": "Html5 Article Engine",
		})
	})

	router.Run(":8080")
}

func loadTemplates() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	adminLayouts, err := fs.Glob(templatePath, "templates/layouts/admin-base.html")
	if err != nil {
		panic(err.Error())
	}

	admins, err := fs.Glob(templatePath, "templates/admins/*.html")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our adminLayouts/ and admins/ directories
	for _, admin := range admins {
		layoutCopy := make([]string, len(adminLayouts))
		copy(layoutCopy, adminLayouts)
		files := append(layoutCopy, admin)
		fmt.Println(files)
		r.AddFromFs(filepath.Base(admin), templatePath, files...)
	}
	return r
}
