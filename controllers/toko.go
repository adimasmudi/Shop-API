package controllers

import (
	"Shop-API/formatter"
	"Shop-API/helper"
	"Shop-API/input"
	"Shop-API/model"
	"Shop-API/services"
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber"
)

type tokoHandler struct {
	tokoService services.ServiceToko
}

func NewTokoHandler(tokoService services.ServiceToko) *tokoHandler{
	return &tokoHandler{tokoService}
}

func (h *tokoHandler) GetMyToko(c *fiber.Ctx){
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	authorizationHeader := c.Get("Authorization")
	tokenString := strings.Split(authorizationHeader," ")[1]

	myToko,user, err := h.tokoService.GetMyToko(tokenString)

	if err != nil{
		response := helper.APIResponse("Failed to get toko data", http.StatusBadRequest, "error", err)
		c.Status(http.StatusBadRequest).JSON(response)
		return
	}

	formatter := formatter.FormatToko(myToko, user)

	response := helper.APIResponse("get toko success", http.StatusOK, "success", formatter)
	
	c.Status(http.StatusOK).JSON(response)
}

func (h *tokoHandler) GetAllToko(c *fiber.Ctx){
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	allToko, err := h.tokoService.GetAllToko()

	if err != nil{
		response := helper.APIResponse("Failed to get all toko data", http.StatusBadRequest, "error", err)
		c.Status(http.StatusBadRequest).JSON(response)
		return
	}

	response := helper.APIResponse("get all toko success", http.StatusOK, "success", allToko)
	
	c.Status(http.StatusOK).JSON(response)
}

func (h *tokoHandler) UpdateToko(c *fiber.Ctx){
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	var input input.UpdateTokoInput

	//validate the request body
	if err := c.BodyParser(&input); err != nil {
		response := helper.APIResponse("Update Toko Failed", http.StatusBadRequest, "error", err)
		c.Status(http.StatusBadRequest).JSON(response)
		return
	}

	params := c.Params("id_toko")

	id, err := strconv.Atoi(params)


	if err != nil{
		response := helper.APIResponse("Update Toko Failed", http.StatusBadRequest, "error", err)
		c.Status(http.StatusBadRequest).JSON(response)
		return
	}

	toko, err := h.tokoService.UpdateToko(id, input)

	if err != nil {
		response := helper.APIResponse("Update Toko Failed", http.StatusBadRequest, "error", err)
		c.Status(http.StatusBadRequest).JSON(response)
		return
	}

	

	formatter := formatter.FormatToko(toko, model.User{})

	response := helper.APIResponse("Toko has been updated", http.StatusOK, "success", formatter)

	c.Status(http.StatusOK).JSON(response)
}