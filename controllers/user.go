package controllers

import (
	"Shop-API/formatter"
	"Shop-API/helper"
	"Shop-API/input"
	"Shop-API/services"
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber"
)

type userHandler struct {
	userService services.Service
}

func NewUserHandler(userService services.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) Register(c *fiber.Ctx) {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	var input input.RegisterUserInput

	//validate the request body
	if err := c.BodyParser(&input); err != nil {
		response := helper.APIResponse("Register Account Failed", http.StatusBadRequest, "error", err)
		c.Status(http.StatusBadRequest).JSON(response)
		return
	}

	newUser, token, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Register Account Failed", http.StatusBadRequest, "error", err)
		c.Status(http.StatusBadRequest).JSON(response)
		return
	}

	formatter := formatter.FormatUser(newUser, token)

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.Status(http.StatusOK).JSON(response)
}

func (h *userHandler) Login(c *fiber.Ctx) {

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	var input input.LoginUserInput

	//validate the request body
	if err := c.BodyParser(&input); err != nil {
		response := helper.APIResponse("Login Failed", http.StatusBadRequest, "error", err)
		c.Status(http.StatusBadRequest).JSON(response)
		return
	}

	loggedinUser, token, err := h.userService.Login(input)

	if err != nil {
		response := helper.APIResponse("Login Failed", http.StatusBadRequest, "error", err)
		c.Status(http.StatusBadRequest).JSON(response)
		return
	}

	formatter := formatter.FormatUser(loggedinUser, token)

	response := helper.APIResponse("successfully logged in", http.StatusOK, "success", formatter)

	c.Status(http.StatusOK).JSON(response)
}

func (h *userHandler) GetProfile(c *fiber.Ctx){

	authorizationHeader := c.Get("Authorization")

	tokenString := strings.Split(authorizationHeader," ")[1]

	
	authUser, err := h.userService.GetProfile(tokenString)

	if err != nil{
		response := helper.APIResponse("Failed to get profile data", http.StatusBadRequest, "error", err)
		c.Status(http.StatusBadRequest).JSON(response)
		return
	}

	formatter := formatter.FormatUser(authUser, tokenString)

	response := helper.APIResponse("get profile success", http.StatusOK, "success", formatter)
	
	c.Status(http.StatusOK).JSON(response)
}

func (h *userHandler) UpdateProfile(c *fiber.Ctx){
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	authorizationHeader := c.Get("Authorization")

	tokenString := strings.Split(authorizationHeader," ")[1]

	user, err := h.userService.GetProfile(tokenString)

	if err != nil{
		response := helper.APIResponse("Can't get current user", http.StatusBadRequest, "error", err)
		c.Status(http.StatusBadRequest).JSON(response)
		return
	}


	var input input.UpdateProfileInput

	//validate the request body
	if err := c.BodyParser(&input); err != nil {
		response := helper.APIResponse("Update Profile Failed", http.StatusBadRequest, "error", err)
		c.Status(http.StatusBadRequest).JSON(response)
		return
	}

	newUser, err := h.userService.UpdateProfile(user.Id,input)

	if err != nil {
		response := helper.APIResponse("Update Profile Failed", http.StatusBadRequest, "error", err)
		c.Status(http.StatusBadRequest).JSON(response)
		return
	}

	formatter := formatter.FormatUser(newUser, tokenString)

	response := helper.APIResponse("Account has been updated", http.StatusOK, "success", formatter)

	c.Status(http.StatusOK).JSON(response)
}