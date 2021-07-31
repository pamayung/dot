package api

import (
	"dot/core/config"
	"dot/core/db"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func init() {
	db.ConnectDB()
}

func TestGetAccount(t *testing.T) {
	// Setup
	e := config.SetupEcho()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/account/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")

	// Assertions
	if assert.NoError(t, GetAccount(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestCreateAccount(t *testing.T) {
	accountJSON := `{"first_name" : "Jon", "last_name" : "Snow", "email" : "jon@gmail.com"}`

	// Setup
	e := config.SetupEcho()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(accountJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, CreateAccount(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUpdateAccount(t *testing.T) {
	accountJSON := `{"first_name" : "Jon", "last_name" : "Snow Update"}`

	// Setup
	e := config.SetupEcho()
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(accountJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/account/:id")
	c.SetParamNames("id")
	c.SetParamValues("11")

	// Assertions
	if assert.NoError(t, UpdateAccount(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDeleteAccount(t *testing.T) {
	// Setup
	e := config.SetupEcho()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/account/:id")
	c.SetParamNames("id")
	c.SetParamValues("11")

	// Assertions
	if assert.NoError(t, DeleteAccount(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
