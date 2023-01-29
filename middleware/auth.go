package middleware

import (
	"os"

	"github.com/gofiber/fiber"
	jwtware "github.com/gofiber/jwt"
)

func Auth() func (c *fiber.Ctx){ 
	return jwtware.New(jwtware.Config{
		ErrorHandler: func (ctx *fiber.Ctx, err error){
			ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error" : "Unauthorized User",
			})
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	})

}