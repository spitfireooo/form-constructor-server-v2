package service_v2

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/request"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/response"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/service"
	"log"
)

func CreateForm(body request.FormWithField, id int) (response.FormWithField, error) {
	form, err := service.CreateForm(request.Form{
		Title:       body.Title,
		Slug:        body.Slug,
		Description: body.Description,
		Logo:        body.Logo,
	}, id)
	if err != nil {
		log.Println("Error in form service", err)
		return response.FormWithField{}, &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: "Error in form service",
		}
	}

	var fields []map[string]interface{}
	if body.Fields != nil {
		for _, field := range *body.Fields {
			if res, err := CreateField(field, int(form.ID)); err != nil {
				log.Println("Error in field service v2", err)
				return response.FormWithField{}, &fiber.Error{
					Code:    fiber.StatusInternalServerError,
					Message: "Error in field service v2",
				}
			} else {
				fields = append(fields, res)
			}
		}
	}

	return response.FormWithField{
		ID:          form.ID,
		Title:       form.Title,
		Slug:        form.Slug,
		Description: form.Description,
		Logo:        form.Logo,
		AuthorId:    form.AuthorId,

		Fields: &fields,
	}, nil
}

func GetForm() {

}
