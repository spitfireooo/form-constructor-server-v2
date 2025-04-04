package service_v2

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/request"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/response"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/service"
	"log"
)

func CreateRadio(body map[string]interface{}, id int) (response.FieldRadio, error) {
	res := new(response.FieldRadio)

	if multiply, err := service.CreateFieldMultiply(request.FieldMultiply{
		IsMultiply: body["is_multiply"].(bool),
	}, id); err != nil {
		return response.FieldRadio{}, err
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
				return response.FieldRadio{}, &fiber.Error{
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

func GetRadio(id int) (response.FieldRadio, error) {
	res := new(response.FieldRadio)

	if multiply, err := service.GetFieldMultiply(id); err != nil {
		return response.FieldRadio{}, err
	} else {
		res.IsMultiply = multiply.IsMultiply
	}

	if fields, err := service.GetFieldVariants(id); err != nil {
		return response.FieldRadio{}, err
	} else {
		res.Variants = fields
	}

	return *res, nil
}

func UpdateRadio(body map[string]interface{}, id int) (response.FieldRadio, error) {
	res := new(response.FieldRadio)

	if multiply, err := service.UpdateFieldMultiply(request.FieldMultiply{
		IsMultiply: body["is_multiply"].(bool),
	}, id); err != nil {
		return response.FieldRadio{}, err
	} else {
		res.IsMultiply = multiply.IsMultiply
	}

	var variantsRes = new([]response.FieldVariants)
	variants := body["variants"].([]interface{})
	if variants != nil {
		for _, variant := range variants {
			Variant := variant.(map[string]interface{})["variant"].(string)
			Name := variant.(map[string]interface{})["name"].(string)

			variantExist, err := service.GetOneFieldVariantsByNameAndFieldID(Name, id)
			if err != nil {
				log.Println("Error in variant service", err)
				return response.FieldRadio{}, &fiber.Error{
					Code:    fiber.StatusInternalServerError,
					Message: "Error in variant service",
				}
			}

			if res, err := service.UpdateFieldVariants(request.FieldVariantsUpdate{
				Variant: &Variant,
				Name:    &Name,
			}, int(variantExist.ID)); err != nil {
				log.Println("Error in variant service", err)
				return response.FieldRadio{}, &fiber.Error{
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

func DeleteRadio(id int) error {
	if err := service.DeleteFieldMultiply(id); err != nil {
		return err
	}

	return nil
}
