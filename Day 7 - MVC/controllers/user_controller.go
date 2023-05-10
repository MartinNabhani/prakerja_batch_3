package controllers

import (
	"myapp/configs"
	"myapp/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	var users []models.User
	if err := configs.DB.Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false, Message: "Gagal get data user dari database", Data: nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true, Message: "Sukses Mendapatkan Semua Users", Data: users,
	})
}

func GetUserByID(c echo.Context) error {

	// ambil data user berdasarkan id
	var user models.User
	if err := configs.DB.First(&user, c.Param("id")).Error; err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Status: false, Message: "Data Tidak DItemukan", Data: nil,
		})
	}

	// return data dalam format JSON
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true, Message: "Data berhasil Ditemukan", Data: user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	if err := configs.DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false, Message: "Gagal Create User", Data: nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true, Message: "Berhasil Create User", Data: user,
	})
}

func DeleteUserController(c echo.Context) error {
	id := c.Param("id")
	var user models.User
	result := configs.DB.First(&user, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{
			Status: false, Message: "Data Tidak Berhasil ditemukan", Data: nil,
		})
	}
	configs.DB.Delete(&user)
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true, Message: "Berhasil Hapus Data", Data: user,
	})
}
func UpdateUserController(c echo.Context) error {
	userID := c.Param("id")

	// Mencari user dengan ID yang diberikan
	var user models.User
	if err := configs.DB.First(&user, userID).Error; err != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{
			Status: false, Message: "Data Tidak Berhasil ditemukan", Data: nil,
		})
	}

	// Mendapatkan data inputan dari client
	var input struct {
		models.User
	}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Memperbarui informasi user
	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password

	if err := configs.DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false, Message: "Gagal Memperbarui data", Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true, Message: "Berhasil Update Data", Data: user,
	})
}
