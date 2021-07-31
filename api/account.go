package api

import (
	"dot/core/db"
	"dot/core/response"
	"dot/model"
	"github.com/labstack/echo/v4"
	"strconv"
)

func GetAccount(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	d := db.DbManager()

	ac := new(model.Account)
	result, err := ac.GetByID(d, uint32(id))
	if err != nil {
		return response.NotFound(c)
	}

	return response.Success(c, result)
}

type updateAccountRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}

func UpdateAccount(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	r := new(updateAccountRequest)
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}

	d := db.DbManager()

	ac := new(model.Account)
	_, err := ac.GetByID(d, uint32(id))
	if err != nil {
		return response.NotFound(c)
	}

	ac.FirstName = r.FirstName
	ac.LastName = r.LastName

	result, err := ac.Update(d, uint32(id))
	if err != nil {
		return response.InternalError(c)
	}

	return response.Success(c, result)
}

type createAccountRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
}

func CreateAccount(c echo.Context) error {
	r := new(createAccountRequest)
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}

	d := db.DbManager()

	ac := new(model.Account)
	ac.FirstName = r.FirstName
	ac.LastName = r.LastName
	ac.Email = r.Email

	mb := new(model.Member)
	mb.Status = 0

	err := ac.Create(d, mb)
	if err != nil {
		return response.InternalError(c)
	}

	return response.Success(c, ac)
}

func DeleteAccount(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	d := db.DbManager()

	ac := new(model.Account)
	_, err := ac.GetByID(d, uint32(id))
	if err != nil {
		return response.NotFound(c)
	}

	_, err = ac.Delete(d, uint32(id))
	if err != nil {
		return response.InternalError(c)
	}

	return response.Success(c, "")
}
