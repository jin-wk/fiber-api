package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jin-wk/fiber-api/config"
	"github.com/jin-wk/fiber-api/database"
	"github.com/jin-wk/fiber-api/models"
	"github.com/jin-wk/fiber-api/utils"
)

// Register godoc
// @Summary     Register
// @Description Register User
// @Tags		Auth
// @Accept		json
// @Produce		json
// @Param		user body models.Register true "user"
// @Router		/api/auth/register [post]
func Register(c *fiber.Ctx) error {
	var register models.Register
	var user struct {
		ID       int    `json:"id"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}

	if err := c.BodyParser(&register); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"message": "Bad Request",
			"data":    err,
		})
	}

	validate := utils.Validate(&register)
	if validate != nil {
		return c.Status(400).JSON(&fiber.Map{
			"message": "Bad Request",
			"data":    validate,
		})
	}

	if database.DB.Table("users").Where("email = ?", register.Email).First(&register).Error == nil {
		return c.Status(409).JSON(&fiber.Map{
			"message": "Email already exists",
			"data":    nil,
		})
	}

	if database.DB.Table("users").Select("email", "password", "name").Create(&register).Error != nil {
		return c.Status(500).JSON(&fiber.Map{
			"message": "Internal server error",
			"data":    nil,
		})
	}

	database.DB.Table("users").Where("email = ?", register.Email).First(&user)
	return c.Status(201).JSON(&fiber.Map{
		"message": "Created",
		"data":    &user,
	})
}

// Login godoc
// @Summary     Login
// @Description Login User
// @Tags		Auth
// @Accept		json
// @Produce		json
// @Param		user body models.Login true "user"
// @Router		/api/auth/login [post]
func Login(c *fiber.Ctx) error {
	var login models.Login
	var user struct {
		ID    int    `json:"id"`
		Email string `json:"email"`
		Name  string `json:"name"`
		Token string `json:"token"`
	}

	if err := c.BodyParser(&login); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"message": "Bad Request",
			"data":    nil,
		})
	}

	validate := utils.Validate(&login)
	if validate != nil {
		return c.Status(400).JSON(&fiber.Map{
			"message": "Bad Request",
			"data":    validate,
		})
	}

	if database.DB.Table("users").Where("email = ?", login.Email).First(&user).Error != nil {
		return c.Status(400).JSON(&fiber.Map{
			"message": "Email does not exists",
			"data":    nil,
		})
	}

	duration, _ := time.ParseDuration(config.Env("JWT_EXPIRE_MIN") + "m")
	claims := jwt.MapClaims{
		"name": user.Name,
		"exp":  time.Now().Add(duration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	user.Token, _ = token.SignedString([]byte(config.Env("JWT_SECRET_KEY")))

	return c.JSON(&fiber.Map{
		"message": "Ok",
		"data":    &user,
	})
}

// Info godoc
// @Summary     Info
// @Description Get Info User
// @Tags		Auth
// @Accept		json
// @Produce		json
// @Param		id path int true "id"
// @Security    Authorization
// @Router		/api/auth/{id} [get]
func Info(c *fiber.Ctx) error {
	var user struct {
		ID        int       `json:"id"`
		Email     string    `json:"email"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	result := database.DB.Table("users").First(&user, c.Params("id"))
	if result.Error != nil {
		return c.Status(404).JSON(&fiber.Map{
			"message": "Not Found",
			"data":    nil,
		})
	}

	return c.JSON(&fiber.Map{
		"message": "Ok",
		"data":    &user,
	})
}
