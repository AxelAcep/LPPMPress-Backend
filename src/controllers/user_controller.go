package controllers

import (
	"lppm/src/database"
	"lppm/src/models"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c echo.Context) error {
	type RegisterInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Nama     string `json:"nama"`
		NoTelp   string `json:"no_telp"`
		Alamat   string `json:"alamat"`
	}

	input := new(RegisterInput)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Format data salah"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal memproses password"})
	}

	tx := database.DB.Begin()

	newUser := models.User{
		Email:    input.Email,
		Password: string(hashedPassword),
		Jenis:    "pelanggan",
	}

	if err := tx.Create(&newUser).Error; err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal buat akun user"})
	}

	newPelanggan := models.Pelanggan{
		UserID:      newUser.ID,
		NamaLengkap: input.Nama,
		NoTelp:      input.NoTelp,
		Alamat:      input.Alamat,
	}

	if err := tx.Create(&newPelanggan).Error; err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal buat profil pelanggan"})
	}

	tx.Commit()

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Registrasi berhasil",
		"user_id": newUser.ID,
	})
}

func RegisterAdmin(c echo.Context) error {
	type RegisterInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Nama     string `json:"nama"`
	}

	input := new(RegisterInput)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Format data salah"})
	}

	tx := database.DB.Begin()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal memproses password"})
	}

	newUser := models.User{
		Email:    input.Email,
		Password: string(hashedPassword),
		Jenis:    "admin",
	}

	if err := tx.Create(&newUser).Error; err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal buat akun user"})
	}

	newAdmin := models.Admin{
		UserID: newUser.ID,
		Nama:   input.Nama,
		User:   newUser,
	}

	if err := tx.Create(&newAdmin).Error; err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal buat profil pelanggan"})
	}

	tx.Commit()

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Registrasi berhasil",
		"user_id": newUser.ID,
	})
}

func LoginUser(c echo.Context) error {
	type LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	input := new(LoginInput)
	c.Bind(input)

	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Email tidak ditemukan"})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Password salah"})
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"jenis":   user.Jenis,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Login berhasil",
		"token":   t,
	})
}
