package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func registerRoutes() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("D:/Go Lang/Golang Gin Framework/templates/*/*.html")
	r.GET("/", func(c *gin.Context) {
		// c.String(http.StatusOK,"Hellow from %v", "Gin")
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.GET("/employees/:id/vacation", func(ctx *gin.Context) {
		id := ctx.Param("id")
		timeOff, ok := TimesOff[id]
		if !ok {
			ctx.String(http.StatusNotFound, "404 - Page Not Found")
			return
		}

		ctx.HTML(http.StatusOK, "vacation-overview.html", map[string]interface{}{
			"TimesOff": timeOff,
		})

	})

	admin := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	}))
	admin.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "admin-overview.html", nil)
	})

	admin.POST("/employees/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "add" {
			pto, err := strconv.ParseFloat(ctx.PostForm("pto"), 32)
			if err != nil {
				ctx.String(http.StatusBadRequest, err.Error())
				return
			}

			startDate, err := time.Parse("2006-01-02", ctx.PostForm("pto"))
			if err != nil {
				ctx.String(http.StatusBadRequest, err.Error())
				return
			}

			var emp Employee
			emp.ID = 42
			emp.FirstName = ctx.PostForm("firstName")
			emp.LastName = ctx.PostForm("lastNAme")
			emp.Position = ctx.PostForm("position")
			emp.Status = "Active"
			emp.TotalPTO = float32(pto)
			emp.StartDate = startDate
			employees["42"] = emp

			ctx.Redirect(http.StatusMovedPermanently, "/admin/employees/42")

		}
	})

	r.Static("/public", "./public")
	// r.StaticFS("/public", http.Dir("./public"))

	return r
}
