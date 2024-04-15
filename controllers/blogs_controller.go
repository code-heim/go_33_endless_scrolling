package controllers

import (
	"fmt"
	"gin_blogs/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BlogsIndex(c *gin.Context) {
	blogs := models.BlogsAll(c)
	content := gin.H{
		"blogs":      blogs,
		"page":       c.GetInt("page"),
		"nextPage":   c.GetInt("page") + 1,
		"pageSize":   c.GetInt("pageSize"),
		"totalPages": c.GetInt("totalPages"),
		"UserID":     c.GetUint("userID"),
	}
	if htmxHeader := c.GetHeader("Hx-Request"); len(htmxHeader) > 0 {
		// This is HTMX request
		if len(*blogs) > 0 {
			c.HTML(
				http.StatusOK,
				"blogs/blog_list.tpl",
				content,
			)
		}
	} else {
		// Regular request
		c.HTML(
			http.StatusOK,
			"blogs/index.tpl",
			content,
		)
	}
}

func BlogsShow(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	blog := models.BlogsFind(id)
	c.HTML(
		http.StatusOK,
		"blogs/show.tpl",
		gin.H{
			"blog":   blog,
			"UserID": c.GetUint("userID"),
		},
	)
}
