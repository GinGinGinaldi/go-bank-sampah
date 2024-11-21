package controllers

import (
	"bank-sampah/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Fungsi untuk menambahkan data sampah
func AddWaste(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var waste models.Waste
		if err := c.Bind(&waste); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		db.Create(&waste)
		return c.JSON(http.StatusCreated, waste)
	}
}

// Fungsi untuk mengambil data sampah berdasarkan User ID
func GetWasteByUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Param("userId")
		var wastes []models.Waste
		db.Where("user_id = ?", userId).Find(&wastes)
		return c.JSON(http.StatusOK, wastes)

	}
}
