package controllers

import (
	"bank-sampah/internal/models"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"github.com/golang-jwt/jwt"
)

var mySigningKey = []byte("secret") // Ganti dengan kunci rahasia yang lebih kuat

// Claims untuk menyimpan informasi token
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Fungsi untuk registrasi pengguna
func Register(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user models.User
		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		// Hash password sebelum disimpan
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Failed to hash password")
		}
		user.Password = string(hashedPassword)

		// Pastikan role diatur sesuai input
		if user.Role == "" {
			user.Role = "warga" // Set default role jika tidak diberikan
		}

		if err := db.Create(&user).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, "Failed to create user")
		}
		return c.JSON(http.StatusCreated, user)
	}
}

// Fungsi untuk login pengguna
func Login(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user models.User
		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		var dbUser models.User
		if err := db.Where("username = ?", user.Username).First(&dbUser).Error; err != nil {
			return c.JSON(http.StatusUnauthorized, "Invalid username or password")
		}

		// Verifikasi password
		if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
			return c.JSON(http.StatusUnauthorized, "Invalid username or password")
		}

		// Buat token
		expirationTime := time.Now().Add(1 * time.Hour)
		claims := &Claims{
			Username: dbUser.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(mySigningKey)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Could not create token")
		}

		// Kirim token sebagai respons
		return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
	}
}

// Middleware untuk memverifikasi token
func TokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, "Missing token")
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return mySigningKey, nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, "Invalid token")
		}

		// Set user information ke konteks
		c.Set("username", claims.Username)
		return next(c)
	}
}

// GetAllUsers mengembalikan daftar semua pengguna
func GetAllUsers(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var users []models.User
		if err := db.Find(&users).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error"})
		}
		if len(users) == 0 {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "No users found"})
		}
		return c.JSON(http.StatusOK, users)
	}
}

// Fungsi untuk mendapatkan statistik
func GetStatistics(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var stats struct {
			TotalVolume     int `json:"total_volume"`
			TotalUsers      int `json:"total_users"`
			TotalCollection int `json:"total_collection"`
		}

		// Menghitung total volume, total pengguna, dan total koleksi
		if err := db.Table("collections").Select("SUM(volume) as total_volume, COUNT(DISTINCT user_id) as total_users, COUNT(collection_id) as total_collection").Scan(&stats).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, "Error retrieving statistics")
		}

		return c.JSON(http.StatusOK, stats)
	}
}

// Fungsi untuk mendapatkan dashboard admin
func GetAdminDashboard(c echo.Context) error {
	// Logika untuk mendapatkan data dashboard admin
	return c.JSON(http.StatusOK, map[string]string{"message": "Welcome to Admin Dashboard!"})
}

// Contoh endpoint yang dilindungi
func GetUserData(c echo.Context) error {
	username := c.Get("username").(string)
	// Logika untuk mendapatkan data pengguna berdasarkan username
	return c.JSON(http.StatusOK, map[string]string{"message": "Welcome " + username})
}