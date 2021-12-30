package middleware

import (
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func LoggerConfig() logger.Config {
	return logger.Config{
		Format:     "${yellow}[${time}] ${blue}[${path}] ${green}[${method}] ${yellow}${body} ${white}${resBody} > ${red}${status}\n",
		TimeFormat: "2006-01-02 15:04:05.000",
		TimeZone:   "Asia/Seoul",
	}
}
