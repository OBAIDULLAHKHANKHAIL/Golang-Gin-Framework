package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var loginCookies = map[string]*loginCookie{}
var identities = []identity{
	{employeeNumber: "1234", password: "password"},
}

const loginCookieName = "Identity"

func loginMiddleware(ctx *gin.Context) {
	if strings.HasPrefix(ctx.Request.URL.Path, "/login") ||
		strings.HasPrefix(ctx.Request.URL.Path, "/public") {
		return
	}

	cookieValue, err := ctx.Cookie(loginCookieName)
	if err != nil {
		ctx.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}

	cookie, ok := loginCookies[cookieValue]

	if !ok ||
		cookie.expiration.Unix() < time.Now().Unix() ||
		cookie.origin != ctx.Request.RemoteAddr {
		ctx.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	ctx.Next()
}

type loginCookie struct {
	value      string
	expiration time.Time
	origin     string
}

type identity struct {
	employeeNumber string
	password       string
}
