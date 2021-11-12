package handlers

import (
	"KVdb/api/v1/dto"
	"KVdb/entity"
	"KVdb/repo"
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type HandErr struct {
	Msg  string `json:"message"`
	Code int    `json:"code"`
}

func GetHandler(rg repo.Getter) echo.HandlerFunc {
	return func(c echo.Context) error {
		key := c.Param("key")
		value, err := rg.Get(context.TODO(), key)
		if err == entity.ErrNotFound {
			return c.JSON(http.StatusNotFound, HandErr{Msg: "Value does not exists", Code: 404})
		} else if err != nil {
			return c.JSON(http.StatusInternalServerError, HandErr{Msg: err.Error(), Code: 500})
		}
		return c.JSON(http.StatusOK, value)
	}
}

func Sethandler(rs repo.Setter) echo.HandlerFunc {
	return func(c echo.Context) error {
		dto := new(dto.SetDto)
		if err := c.Bind(&dto); err != nil {
			return c.JSON(http.StatusBadRequest, HandErr{Msg: err.Error(), Code: 400})
		}
		if err := dto.Validate(); err != nil {
			return c.JSON(http.StatusBadRequest, HandErr{Msg: err.Error(), Code: 400})
		}

		exp := time.Now().Add(time.Duration(dto.Exp) * time.Second)
		if err := rs.Set(context.TODO(), dto.Key, dto.Value, exp); err != nil {
			return c.JSON(http.StatusInternalServerError, HandErr{Msg: err.Error(), Code: 500})
		}
		return c.NoContent(http.StatusOK)
	}
}
