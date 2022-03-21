package controller

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
// @Accept		application/json
// @Produce		application/json
// @Param		user body models.RegisterUser true "user"
// @Success		201 {object} utils.Resp
// @Failure		500 {object} utils.Resp
// @Router		/api/auth/register [post]
func Register(c *fiber.Ctx) error {
	user := new(models.RegisterUser)

	if err := c.BodyParser(&user); err != nil {
		return utils.Response(c, 400, "Bad Request", err)
	}

	err := utils.Validate(user)
	if err != nil {
		return utils.Response(c, 400, "Bad Request", err)
	}

	if database.DB.Model(user).Where("email = ?", user.Email).First(&user).Error == nil {
		return utils.Response(c, 409, "Email Already Exists", nil)
	}

	if database.DB.Select("email", "password", "name").Create(&user).Error != nil {
		return utils.Response(c, 500, "Internal Server Error", nil)
	}

	return utils.Response(c, 201, "Created", &user)
}

// Login godoc
// @Summary     Login
// @Description Login User
// @Tags		Auth
// @Accept		application/json
// @Produce		application/json
// @Param		user body models.LoginUser true "user"
// @Success		200 {object} utils.Resp
// @Failure		500 {object} utils.Resp
// @Router		/api/auth/login [post]
func Login(c *fiber.Ctx) error {
	var loginUser models.LoginUser
	var responseUser models.ResponseUser

	if err := c.BodyParser(&loginUser); err != nil {
		return utils.Response(c, 400, "Bad Request", err)
	}

	if database.DB.Model(&loginUser).Where("email = ?", loginUser.Email).First(&responseUser).Error != nil {
		return utils.Response(c, 401, "Email Not Exists", nil)
	}

	duration, _ := time.ParseDuration(config.Env("JWT_EXPIRE_MIN") + "m")
	claims := jwt.MapClaims{
		"name": responseUser.Name,
		"exp":  time.Now().Add(duration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.Env("JWT_SECRET_KEY")))
	if err != nil {
		return utils.Response(c, 500, "Internal Server Error", nil)
	}

	return utils.Response(c, 200, "OK", map[string]interface{}{
		"ID":    responseUser.ID,
		"Email": responseUser.Email,
		"Name":  responseUser.Name,
		"token": t,
	})
}

// Info godoc
// @Summary     Info
// @Description Get Info User
// @Tags		Auth
// @Accept		application/json
// @Produce		application/json
// @Param		id path int true "id"
// @Success		200 {object} utils.Resp
// @Failure		404 {object} utils.Resp
// @Security    Authorization
// @Router		/api/auth/{id} [get]
func Info(c *fiber.Ctx) error {
	var user models.ResponseUser

	result := database.DB.First(&user, c.Params("id"))
	if result.Error != nil {
		return utils.Response(c, 404, "Not Found", nil)
	}

	return utils.Response(c, 200, "OK", &user)
}
