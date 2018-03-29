/*
 * Copyright 2018 Foolin.  All rights reserved.
 *
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/foolin/gin-template"
)

func main() {

	router := gin.Default()

	//new template engine
	router.HTMLRender = gintemplate.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index", gin.H{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
		})
	})


	router.GET("/block", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "block", gin.H{"title": "Block file title!!"})
	})

	router.Run(":9090")
}
