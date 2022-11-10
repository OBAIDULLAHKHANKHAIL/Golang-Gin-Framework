package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func registerRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(loginMiddleware)

	r.LoadHTMLGlob("D:/Go Lang/Golang Gin Framework/templates/*/*.html")
	r.Any("/", func(ctx  *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	r.Any("/login", func(ctx  *gin.Context) {
		employeeNumber := ctx.PostForm("employeeNumber")
		password := ctx.PostForm("password")

		for _, identity := range identities {
			if identity.employeeNumber == employeeNumber &&
				identity.password == password {
				lc := loginCookie{
					value:      employeeNumber,
					expiration: time.Now().Add(24 * time.Hour),
					origin:     ctx.Request.RemoteAddr,
				}

				loginCookies[lc.value] = &lc

				maxAge := lc.expiration.Unix() - time.Now().Unix()
				ctx.SetCookie(loginCookieName, lc.value, int(maxAge), "", "", false, true)

				ctx.Redirect(http.StatusTemporaryRedirect, "/")
				return
			}
		}

		ctx.HTML(http.StatusOK, "login.html", nil)
	})

	r.GET("/employees/:id/vacation", func(ctx  *gin.Context) {
		id := ctx.Param("id")
		timesOff, ok := TimesOff[id]

		if !ok {
			ctx.String(http.StatusNotFound, "404 - Page Not Found")
			return
		}

		ctx.HTML(http.StatusOK, "vacation-overview.html",
			gin.H{
				"TimesOff": timesOff,
			})
	})

	r.POST("/employees/:id/vacation/new", func(ctx  *gin.Context) {
		var timeOff TimeOff
		err := ctx.BindJSON(&timeOff)

		if err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			return
		}

		id := ctx.Param("id")

		timesOff, ok := TimesOff[id]

		if !ok {
			TimesOff[id] = []TimeOff{}
		}

		TimesOff[id] = append(timesOff, timeOff)

		ctx.JSON(http.StatusCreated, &timeOff)
	})

	admin := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	}))
	admin.GET("/", func(ctx  *gin.Context) {
		ctx.HTML(http.StatusOK, "admin-overview.html",
			gin.H{
				"Employees": employees,
			})
	})

	admin.GET("/employees/:id", func(ctx  *gin.Context) {
		id := ctx.Param("id")
		if id == "add" {
			ctx.HTML(http.StatusOK, "admin-employee-add.html", nil)
			return
		}
		employee, ok := employees[id]

		if !ok {
			ctx.String(http.StatusNotFound, "404 - Not Found")
			return
		}

		ctx.HTML(http.StatusOK, "admin-employee-edit.html",
			gin.H{
				"Employee": employee,
			})
	})

	admin.POST("/employees/:id", func(ctx  *gin.Context) {
		id := ctx.Param("id")

		if id == "add" {

			startDate, err := time.Parse("2006-01-02", ctx.PostForm("startDate"))
			if err != nil {
				ctx.String(http.StatusBadRequest, err.Error())
				return
			}

			var emp Employee
			ctx.Bind(&emp)
			emp.ID = 42
			emp.Status = "Active"
			emp.StartDate = startDate
			employees["42"] = emp

			ctx.Redirect(http.StatusMovedPermanently, "/admin/employees/42")
		}
	})

	r.Static("/public", "./public")

	return r
}
