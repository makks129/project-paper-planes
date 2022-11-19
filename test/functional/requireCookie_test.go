package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/makks129/project-paper-planes/src/router"
	"github.com/makks129/project-paper-planes/test/suit"
	"github.com/stretchr/testify/assert"
)

const COOKIE_TEST = "cookie_test"

func Test_RequireCookie(t *testing.T) {
	app := InitApp()
	app.GET("/test", router.RequireCookie(COOKIE_TEST))

	s := suit.Of(&suit.SubTests{T: t})

	s.Test("returns 200, if cookie exists", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		req.AddCookie(&http.Cookie{Name: COOKIE_TEST, Value: ALICE_ID, Secure: true, HttpOnly: true})
		app.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})

	s.Test("returns 400, if cookie does not exist", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		// no cookie
		app.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	})
}
