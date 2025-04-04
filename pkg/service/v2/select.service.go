package service_v2

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/request"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/response"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/service"
	"log"
)

func CreateSelect(body map[string]interface{}, id int) (response.FieldSelect, error) {
	res := new(response.FieldSelect)

	if placeholder, err := service.CreateFieldPlaceholder(request.FieldPlaceholder{
		Placeholder: body["placeholder"].(string),
	}, id); err != nil {
		return response.FieldSelect{}, err
	} else {
		res.Placeholder = placeholder.Placeholder
	}

	if multiply, err := service.CreateFieldMultiply(request.FieldMultiply{
		IsMultiply: body["is_multiply"].(bool),
	}, id); err != nil {
		return response.FieldSelect{}, err
	} else {
		res.IsMultiply = multiply.IsMultiply
	}

	var variantsRes = new([]response.FieldVariants)
	variants := body["variants"].([]interface{})
	if variants != nil {
		for _, variant := range variants {
			if res, err := service.CreateFieldVariant(request.FieldVariants{
				Variant: variant.(map[string]interface{})["variant"].(string),
				Name:    variant.(map[string]interface{})["name"].(string),
			}, id); err != nil {
				log.Println("Error in variant service", err)
				return response.FieldSelect{}, &fiber.Error{
					Code:    fiber.StatusInternalServerError,
					Message: "Error in variant service",
				}
			} else {
				*variantsRes = append(*variantsRes, res)
			}
		}
	}
	res.Variants = *variantsRes

	return *res, nil
}

func GetSelect(id int) (response.FieldSelect, error) {
	res := new(response.FieldSelect)

	if placeholder, err := service.GetFieldPlaceholder(id); err != nil {
		return response.FieldSelect{}, err
	} else {
		res.Placeholder = placeholder.Placeholder
	}

	if multiply, err := service.GetFieldMultiply(id); err != nil {
		return response.FieldSelect{}, err
	} else {
		res.IsMultiply = multiply.IsMultiply
	}

	if fields, err := service.GetFieldVariants(id); err != nil {
		return response.FieldSelect{}, err
	} else {
		res.Variants = fields
	}

	return *res, nil
}

func UpdateSelect(body map[string]interface{}, id int) (response.FieldSelect, error) {
	res := new(response.FieldSelect)

	if placeholder, err := service.UpdateFieldPlaceholder(request.FieldPlaceholder{
		Placeholder: body["placeholder"].(string),
	}, id); err != nil {
		return response.FieldSelect{}, err
	} else {
		res.Placeholder = placeholder.Placeholder
	}

	if multiply, err := service.UpdateFieldMultiply(request.FieldMultiply{
		IsMultiply: body["is_multiply"].(bool),
	}, id); err != nil {
		return response.FieldSelect{}, err
	} else {
		res.IsMultiply = multiply.IsMultiply
	}

	return *res, nil
}

func DeleteSelect(id int) error {
	if err := service.DeleteFieldPlaceholder(id); err != nil {
		return err
	}

	if err := service.DeleteFieldMultiply(id); err != nil {
		return err
	}

	return nil
}
