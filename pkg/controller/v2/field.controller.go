package controller_v2

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/request"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/service"
	"log"
	"net/http"
)

func CreateField(ctx *fiber.Ctx) error {
	formId, _ := ctx.ParamsInt("formId")

	var body map[string]interface{}
	if err := ctx.BodyParser(&body); err != nil {
		log.Println("Error in parsing request", err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in parsing request",
		})
	}

	field, err := service.CreateField(request.Field{
		Name:     body["name"].(string),
		Type:     body["type"].(string),
		Label:    body["label"].(string),
		OrderOf:  int(body["order_of"].(float64)),
		Required: body["required"].(bool),
	}, formId)
	if err != nil {
		log.Println("Error in field service", err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error in field service",
		})
	}

	resp := map[string]interface{}{
		"id":       field.ID,
		"name":     field.Name,
		"type":     field.Type,
		"label":    field.Label,
		"order_od": field.OrderOf,
		"required": field.Required,
	}

	if t, ok := body["type"]; ok {
		switch t {
		case "string":
			if res, err := service.CreateString(body, int(field.ID)); err != nil {
				log.Println("Error in string service", err)
				return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
					"message": "Error in string service",
				})
			} else {
				resp["placeholder"] = res.Placeholder
				resp["min"] = res.Min
				resp["max"] = res.Max
			}

		case "number":
			if res, err := service.CreateNumber(body, int(field.ID)); err != nil {
				log.Println("Error in number service", err)
				return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
					"message": "Error in number service",
				})
			} else {
				resp["placeholder"] = res.Placeholder
				resp["min"] = res.Min
				resp["max"] = res.Max
			}

		case "email":
			if res, err := service.CreateEmail(body, int(field.ID)); err != nil {
				log.Println("Error in email service", err)
				return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
					"message": "Error in email service",
				})
			} else {
				resp["placeholder"] = res.Placeholder
				resp["min"] = res.Min
				resp["max"] = res.Max
			}

		case "text":
			if res, err := service.CreateText(body, int(field.ID)); err != nil {
				log.Println("Error in text service", err)
				return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
					"message": "Error in text service",
				})
			} else {
				resp["placeholder"] = res.Placeholder
				resp["min"] = res.Min
				resp["max"] = res.Max
			}

		case "date":
			if res, err := service.CreateDate(body, int(field.ID)); err != nil {
				log.Println("Error in date service", err)
				return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
					"message": "Error in date service",
				})
			} else {
				resp["placeholder"] = res.Placeholder
				resp["min"] = res.Min
				resp["max"] = res.Max
			}

		case "radio":
			if res, err := service.CreateRadio(body, int(field.ID)); err != nil {
				log.Println("Error in radio service", err)
				return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
					"message": "Error in radio service",
				})
			} else {
				resp["is_multiply"] = res.IsMultiply
			}

		case "select":
			if res, err := service.CreateSelect(body, int(field.ID)); err != nil {
				log.Println("Error in select service", err)
				return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
					"message": "Error in select service",
				})
			} else {
				resp["placeholder"] = res.Placeholder
				resp["is_multiply"] = res.IsMultiply
			}
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": resp,
	})
}

func GetOneField(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	field, err := service.GetOneField(id)
	if err != nil {
		log.Println("Error in field service", err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error in field service",
		})
	}

	resp := map[string]interface{}{
		"id":       field.ID,
		"name":     field.Name,
		"type":     field.Type,
		"label":    field.Label,
		"order_od": field.OrderOf,
		"required": field.Required,
	}

	switch field.Type {
	case "string":
		if res, err := service.GetString(id); err != nil {
			log.Println("Error in string service", err)
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error in string service",
			})
		} else {
			resp["placeholder"] = res.Placeholder
			resp["min"] = res.Min
			resp["max"] = res.Max
		}

	case "number":
		if res, err := service.GetNumber(id); err != nil {
			log.Println("Error in number service", err)
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error in number service",
			})
		} else {
			resp["placeholder"] = res.Placeholder
			resp["min"] = res.Min
			resp["max"] = res.Max
		}

	case "email":
		if res, err := service.GetEmail(id); err != nil {
			log.Println("Error in email service", err)
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error in email service",
			})
		} else {
			resp["placeholder"] = res.Placeholder
			resp["min"] = res.Min
			resp["max"] = res.Max
		}

	case "text":
		if res, err := service.GetText(id); err != nil {
			log.Println("Error in text service", err)
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error in text service",
			})
		} else {
			resp["placeholder"] = res.Placeholder
			resp["min"] = res.Min
			resp["max"] = res.Max
		}

	case "date":
		if res, err := service.GetDate(id); err != nil {
			log.Println("Error in date service", err)
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error in date service",
			})
		} else {
			resp["placeholder"] = res.Placeholder
			resp["min"] = res.Min
			resp["max"] = res.Max
		}

	case "radio":
		if res, err := service.GetRadio(id); err != nil {
			log.Println("Error in radio service", err)
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error in radio service",
			})
		} else {
			resp["is_multiply"] = res.IsMultiply
		}

	case "select":
		if res, err := service.GetSelect(id); err != nil {
			log.Println("Error in select service", err)
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error in select service",
			})
		} else {
			resp["placeholder"] = res.Placeholder
			resp["is_multiply"] = res.IsMultiply
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": resp,
	})
}

func UpdateField(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "UpdateField",
	})
}
