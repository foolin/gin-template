package main

import (
	"github.com/gin-gonic/gin"
	"github.com/foolin/gin-template"
	"net/http"
	"html/template"
)

func main() {
	router := gin.Default()

	//new template engine
	router.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:      "views",
		Extension: ".tpl",
		Master:    "layouts/master",
		Partials:  []string{"partials/head"},
		Funcs: template.FuncMap{
			"sub": func(a, b int) int {
				return a - b
			},
		},
		DisableCache: false,
	})

	router.GET("/", func(ctx *gin.Context) {
		//render with master
		ctx.HTML(http.StatusOK, "index", gin.H{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
		})
	})

	router.GET("/page_file", func(ctx *gin.Context) {
		//render only file, must full name with extension
		ctx.HTML(http.StatusOK, "page_file.tpl", gin.H{"title": "Page file title!!"})
	})

	router.Run(":9090")
}
