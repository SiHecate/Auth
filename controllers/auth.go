package controllers

import (
	"Auth/database"
	"Auth/models"
	"Auth/utils"
	"strconv"
	"strings"
	"time"

	"Auth/helpers"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *fiber.Ctx) error {
	var signupRequest struct {
		Name             string `json:"name"`
		Lastname         string `json:"lastname"`
		Username         string `json:"username"`
		Gender           string `json:"gender"`
		Email            string `json:"email"`
		Password         string `json:"password"`
		Password_confirm string `json:"password_confirm"`
	}

	signupRequest.Name = strings.TrimSpace(signupRequest.Name)
	signupRequest.Lastname = strings.TrimSpace(signupRequest.Lastname)
	signupRequest.Username = strings.TrimSpace(signupRequest.Username)
	signupRequest.Email = strings.TrimSpace(signupRequest.Email)

	if err := c.BodyParser(&signupRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request data: " + err.Error()})
	}

	if signupRequest.Password != signupRequest.Password_confirm {
		return c.Status(400).JSON(fiber.Map{"error": "Passwords do not match"})
	}

	if helpers.ValidEmail(signupRequest.Email) {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid email format"})
	}

	if helpers.StrongPassword(signupRequest.Password) {
		return c.Status(400).JSON(fiber.Map{"error": "Password is not strong enough"})
	}

	var existingUser models.User
	database.Conn.Where("email = ?", signupRequest.Email).First(&existingUser)

	if existingUser.ID != 0 {
		return c.Status(400).JSON(fiber.Map{"error": "user already exists"})
	}

	if existingUser.Username != "" {
		return c.Status(400).JSON(fiber.Map{"error": "username already exists"})
	}

	if signupRequest.Gender != "Woman" && signupRequest.Gender != "Man" && signupRequest.Gender != "Other" {
		return c.Status(400).JSON(fiber.Map{"error": "Gender must be either 'Woman','Man' or `Other`."})
	}

	var err error
	signupRequest.Password, err = utils.GenerateHashPassword(signupRequest.Password)
	if err != nil {
		return err
	}

	response := models.User{
		Name:     signupRequest.Name,
		Lastname: signupRequest.Lastname,
		Username: signupRequest.Username,
		Gender:   signupRequest.Gender,
		Email:    signupRequest.Email,
		Password: signupRequest.Password,
		Role:     "User",
	}

	if err := database.Conn.Create(&response).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "User creation failed"})
	}

	return c.JSON(response)
}

func Login(c *fiber.Ctx) error {
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&loginRequest); err != nil {
		c.Status(400).JSON(fiber.Map{"error": err.Error()})
		return err
	}

	var existingUser models.User

	if err := database.Conn.Where("email = ?", loginRequest.Email).First(&existingUser).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid email",
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(loginRequest.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid password",
		})
	}

	expirationTime := time.Now().Add(60 * time.Minute)

	claims := &models.Claims{
		Role: existingUser.Role,
		StandardClaims: jwt.StandardClaims{
			Issuer:    strconv.Itoa(int(existingUser.ID)),
			Subject:   existingUser.Email,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	var jwtKey = []byte("my_secret_key")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.Status(500).JSON(fiber.Map{"error": "could not generate token"})
		return err
	}

	cookie := fiber.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Minute * 60),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "user login success!",
	})
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
