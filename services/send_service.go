package services

import (
	"github.com/aldinokemal/go-whatsapp-web-multidevice/structs"
	"github.com/gofiber/fiber/v2"
)

type SendService interface {
	SendText(c *fiber.Ctx, request structs.SendMessageRequest) (response structs.SendMessageResponse, err error)
	SendImage(c *fiber.Ctx)
}
