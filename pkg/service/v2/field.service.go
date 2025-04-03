package service_v2

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/request"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/response"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/service"
	"log"
)

type FieldResponse interface {
	response.FieldString |
		response.FieldNumber |
		response.FieldEmail |
		response.FieldText |
		response.FieldDate |
		response.FieldRadio |
		response.FieldSelect
}

type FieldRequest interface {
	request.FieldString |
		request.FieldNumber |
		request.FieldEmail |
		request.FieldText |
		request.FieldDate |
		request.FieldRadio |
		request.FieldSelect
}

func CreateField(body map[string]interface{}, formId int) (map[string]interface{}, error) {
	field, err := service.CreateField(request.Field{
		Name:     body["name"].(string),
		Type:     body["type"].(string),
		Label:    body["label"].(string),
		OrderOf:  int(body["order_of"].(float64)),
		Required: body["required"].(bool),
	}, formId)
	if err != nil {
		log.Println("Error in field service", err)
		return nil, &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: "Error in field service",
		}
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
			if res, err := CreateString(body, int(field.ID)); err != nil {
				log.Println("Error in string service", err)
				return nil, &fiber.Error{
					Code:    fiber.StatusInternalServerError,
					Message: "Error in string service",
				}
			} else {
				resp["placeholder"] = res.Placeholder
				resp["min"] = res.Min
				resp["max"] = res.Max
			}

		case "number":
			if res, err := CreateNumber(body, int(field.ID)); err != nil {
				log.Println("Error in number service", err)
				return nil, &fiber.Error{
					Code:    fiber.StatusInternalServerError,
					Message: "Error in number service",
				}
			} else {
				resp["placeholder"] = res.Placeholder
				resp["min"] = res.Min
				resp["max"] = res.Max
			}

		case "email":
			if res, err := CreateEmail(body, int(field.ID)); err != nil {
				log.Println("Error in email service", err)
				return nil, &fiber.Error{
					Code:    fiber.StatusInternalServerError,
					Message: "Error in email service",
				}
			} else {
				resp["placeholder"] = res.Placeholder
				resp["min"] = res.Min
				resp["max"] = res.Max
			}

		case "text":
			if res, err := CreateText(body, int(field.ID)); err != nil {
				log.Println("Error in text service", err)
				return nil, &fiber.Error{
					Code:    fiber.StatusInternalServerError,
					Message: "Error in text service",
				}
			} else {
				resp["placeholder"] = res.Placeholder
				resp["min"] = res.Min
				resp["max"] = res.Max
			}

		case "date":
			if res, err := CreateDate(body, int(field.ID)); err != nil {
				log.Println("Error in date service", err)
				return nil, &fiber.Error{
					Code:    fiber.StatusInternalServerError,
					Message: "Error in date service",
				}
			} else {
				resp["placeholder"] = res.Placeholder
				resp["min"] = res.Min
				resp["max"] = res.Max
			}

		case "radio":
			if res, err := CreateRadio(body, int(field.ID)); err != nil {
				log.Println("Error in radio service", err)
				return nil, &fiber.Error{
					Code:    fiber.StatusInternalServerError,
					Message: "Error in radio service",
				}
			} else {
				resp["is_multiply"] = res.IsMultiply
			}

		case "select":
			if res, err := CreateSelect(body, int(field.ID)); err != nil {
				log.Println("Error in select service", err)
				return nil, &fiber.Error{
					Code:    fiber.StatusInternalServerError,
					Message: "Error in select service",
				}
			} else {
				resp["placeholder"] = res.Placeholder
				resp["is_multiply"] = res.IsMultiply
			}
		}
	}

	return resp, nil
}

func GetOneField(id int) (map[string]interface{}, error) {
	field, err := service.GetOneField(id)
	if err != nil {
		log.Println("Error in field service", err)
		return nil, &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: "Error in field service",
		}
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
		if res, err := GetString(id); err != nil {
			log.Println("Error in string service", err)
			return nil, &fiber.Error{
				Code:    fiber.StatusInternalServerError,
				Message: "Error in string service",
			}
		} else {
			resp["placeholder"] = res.Placeholder
			resp["min"] = res.Min
			resp["max"] = res.Max
		}

	case "number":
		if res, err := GetNumber(id); err != nil {
			log.Println("Error in number service", err)
			return nil, &fiber.Error{
				Code:    fiber.StatusInternalServerError,
				Message: "Error in number service",
			}
		} else {
			resp["placeholder"] = res.Placeholder
			resp["min"] = res.Min
			resp["max"] = res.Max
		}

	case "email":
		if res, err := GetEmail(id); err != nil {
			log.Println("Error in email service", err)
			return nil, &fiber.Error{
				Code:    fiber.StatusInternalServerError,
				Message: "Error in email service",
			}
		} else {
			resp["placeholder"] = res.Placeholder
			resp["min"] = res.Min
			resp["max"] = res.Max
		}

	case "text":
		if res, err := GetText(id); err != nil {
			log.Println("Error in text service", err)
			return nil, &fiber.Error{
				Code:    fiber.StatusInternalServerError,
				Message: "Error in text service",
			}
		} else {
			resp["placeholder"] = res.Placeholder
			resp["min"] = res.Min
			resp["max"] = res.Max
		}

	case "date":
		if res, err := GetDate(id); err != nil {
			log.Println("Error in date service", err)
			return nil, &fiber.Error{
				Code:    fiber.StatusInternalServerError,
				Message: "Error in date service",
			}
		} else {
			resp["placeholder"] = res.Placeholder
			resp["min"] = res.Min
			resp["max"] = res.Max
		}

	case "radio":
		if res, err := GetRadio(id); err != nil {
			log.Println("Error in radio service", err)
			return nil, &fiber.Error{
				Code:    fiber.StatusInternalServerError,
				Message: "Error in radio service",
			}
		} else {
			resp["is_multiply"] = res.IsMultiply
		}

	case "select":
		if res, err := GetSelect(id); err != nil {
			log.Println("Error in select service", err)
			return nil, &fiber.Error{
				Code:    fiber.StatusInternalServerError,
				Message: "Error in select service",
			}
		} else {
			resp["placeholder"] = res.Placeholder
			resp["is_multiply"] = res.IsMultiply
		}
	}

	return resp, err
}

func UpdateField(body map[string]interface{}, id int) (map[string]interface{}, error) {
	oldField, err := service.GetOneField(id)
	if err != nil {
		log.Println("Error in field service", err)
		return nil, &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: "Error in field service",
		}
	}

	Name := body["name"].(string)
	Type := body["type"].(string)
	Label := body["label"].(string)
	OrderOf := int(body["order_of"].(float64))
	Required := body["required"].(bool)

	field, err := service.UpdateField(request.FieldUpdate{
		Name:     &Name,
		Type:     &Type,
		Label:    &Label,
		OrderOf:  &OrderOf,
		Required: &Required,
	}, id)
	if err != nil {
		log.Println("Error in field service", err)
		return nil, &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: "Error in field service",
		}
	}

	resp := map[string]interface{}{
		"id":       field.ID,
		"name":     field.Name,
		"type":     field.Type,
		"label":    field.Label,
		"order_od": field.OrderOf,
		"required": field.Required,
	}

	// ...
	if field.Type != "" && oldField.Type != field.Type {
		fieldID := int(field.ID)
		switch oldField.Type {
		case "string":
			DeleteString(fieldID)
		case "number":
			DeleteNumber(fieldID)

		case "email":
			DeleteEmail(fieldID)

		case "text":
			DeleteText(fieldID)

		case "date":
			DeleteDate(fieldID)

		case "radio":
			DeleteRadio(fieldID)

		case "select":
			DeleteSelect(fieldID)
		}
	}

	if t, ok := body["type"]; ok {
		switch t {
		case "string":
			var res response.FieldString
			if oldField.Type == field.Type {
				if res, err = UpdateString(body, int(field.ID)); err != nil {
					log.Println("Error in string service", err)
					return nil, &fiber.Error{
						Code:    fiber.StatusInternalServerError,
						Message: "Error in string service",
					}
				}
			} else {
				if res, err = CreateString(body, int(field.ID)); err != nil {
					log.Println("Error in string service", err)
					return nil, &fiber.Error{
						Code:    fiber.StatusInternalServerError,
						Message: "Error in string service",
					}
				}
			}
			resp["placeholder"] = res.Placeholder
			resp["min"] = res.Min
			resp["max"] = res.Max

		case "number":
			var res response.FieldNumber
			if oldField.Type == field.Type {
				if res, err = UpdateNumber(body, int(field.ID)); err != nil {
					log.Println("Error in number service", err)
					return nil, &fiber.Error{
						Code:    fiber.StatusInternalServerError,
						Message: "Error in number service",
					}
				}
			} else {
				if res, err = CreateNumber(body, int(field.ID)); err != nil {
					log.Println("Error in number service", err)
					return nil, &fiber.Error{
						Code:    fiber.StatusInternalServerError,
						Message: "Error in number service",
					}
				}
			}
			resp["placeholder"] = res.Placeholder
			resp["min"] = res.Min
			resp["max"] = res.Max

		case "email":
			var res response.FieldEmail
			if oldField.Type == field.Type {
				if res, err = UpdateEmail(body, int(field.ID)); err != nil {
					log.Println("Error in email service", err)
					return nil, &fiber.Error{
						Code:    fiber.StatusInternalServerError,
						Message: "Error in email service",
					}
				}
			} else {
				if res, err = CreateEmail(body, int(field.ID)); err != nil {
					log.Println("Error in email service", err)
					return nil, &fiber.Error{
						Code:    fiber.StatusInternalServerError,
						Message: "Error in email service",
					}
				}
			}
			resp["placeholder"] = res.Placeholder
			resp["min"] = res.Min
			resp["max"] = res.Max

		case "text":
			var res response.FieldText
			if oldField.Type == field.Type {
				if res, err = UpdateText(body, int(field.ID)); err != nil {
					log.Println("Error in text service", err)
					return nil, &fiber.Error{
						Code:    fiber.StatusInternalServerError,
						Message: "Error in text service",
					}
				}
			} else {
				if res, err = CreateText(body, int(field.ID)); err != nil {
					log.Println("Error in text service", err)
					return nil, &fiber.Error{
						Code:    fiber.StatusInternalServerError,
						Message: "Error in text service",
					}
				}
			}
			resp["placeholder"] = res.Placeholder
			resp["min"] = res.Min
			resp["max"] = res.Max

		case "date":
			var res response.FieldDate
			if oldField.Type == field.Type {
				if res, err = UpdateDate(body, int(field.ID)); err != nil {
					log.Println("Error in date service", err)
					return nil, &fiber.Error{
						Code:    fiber.StatusInternalServerError,
						Message: "Error in date service",
					}
				}
			} else {
				if res, err = CreateDate(body, int(field.ID)); err != nil {
					log.Println("Error in date service", err)
					return nil, &fiber.Error{
						Code:    fiber.StatusInternalServerError,
						Message: "Error in date service",
					}
				}
			}
			resp["placeholder"] = res.Placeholder
			resp["min"] = res.Min
			resp["max"] = res.Max

		case "radio":
			var res response.FieldRadio
			if oldField.Type == field.Type {
				if res, err = UpdateRadio(body, int(field.ID)); err != nil {
					log.Println("Error in radio service", err)
					return nil, &fiber.Error{
						Code:    fiber.StatusInternalServerError,
						Message: "Error in radio service",
					}
				}
			} else {
				if res, err = CreateRadio(body, int(field.ID)); err != nil {
					log.Println("Error in radio service", err)
					return nil, &fiber.Error{
						Code:    fiber.StatusInternalServerError,
						Message: "Error in radio service",
					}
				}
			}
			resp["is_multiply"] = res.IsMultiply

		case "select":
			var res response.FieldSelect
			if oldField.Type == field.Type {
				if res, err = UpdateSelect(body, int(field.ID)); err != nil {
					log.Println("Error in select service", err)
					return nil, &fiber.Error{
						Code:    fiber.StatusInternalServerError,
						Message: "Error in select service",
					}
				}
			} else {
				if res, err = CreateSelect(body, int(field.ID)); err != nil {
					log.Println("Error in select service", err)
					return nil, &fiber.Error{
						Code:    fiber.StatusInternalServerError,
						Message: "Error in select service",
					}
				}
			}
			resp["placeholder"] = res.Placeholder
			resp["is_multiply"] = res.IsMultiply
		}
	}

	return resp, nil
}
