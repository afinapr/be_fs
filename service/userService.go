package service

import (
	"be_fs/models"
	"be_fs/repository"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type ResponseModel struct {
	Code    int    `json: "code" validate: "required"`
	Message string `json: "message" validate: "required"`
}

// Createfunction
func CreateUser(c echo.Context) error {
	Res := &ResponseModel{400, "Bad Request"}
	U := new(models.User)
	if err := c.Bind(U); err != nil {
		return nil
	}
	Res = (*ResponseModel)(repository.CreateUser(U))
	return c.JSON(http.StatusOK, Res)
}

// ReadAllUserFunction
func ReadAllUser(c echo.Context) error {
	limit := c.Param("limit")
	dataLimit, _ := strconv.Atoi(limit)
	offset := c.Param("offset")
	dataOffset, _ := strconv.Atoi(offset)

	result := repository.ReadAllUser(dataLimit, dataOffset)
	return c.JSON(http.StatusOK, result)
}

// ReadAllUserFunction
func ReadAllUserNoLimit(c echo.Context) error {
	result := repository.ReadAllUserNoLimit()
	return c.JSON(http.StatusOK, result)
}

// ReadIdUserFunction
func ReadIdUser(c echo.Context) error {
	id := c.Param("userId")
	data, _ := strconv.Atoi(id)
	result := repository.ReadByIdUser(data)
	return c.JSON(http.StatusOK, result)
}

// UpdateFunction
func UpdateUser(c echo.Context) error {
	Res := &ResponseModel{400, "Bad Request"}
	id := c.Param("userId")
	data, _ := strconv.Atoi(id)
	U := new(models.User)
	if err := c.Bind(U); err != nil {
		log.Println(err.Error())
		return nil
	}
	Res = (*ResponseModel)(repository.UpdateUser(U, data))
	return c.JSON(http.StatusOK, Res)
}

// DeleteFunction
func DeleteUser(c echo.Context) error {
	Res := &ResponseModel{400, "Bad Request"}
	id := c.Param("userId")
	data, _ := strconv.Atoi(id)
	fmt.Println("userId", data)
	Res = (*ResponseModel)(repository.DeleteUser(data))
	return c.JSON(http.StatusOK, Res)
}
