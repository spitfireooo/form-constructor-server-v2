package controller_v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/request"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/response"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/service"
	"log"
	"net/http"
)

// @Summary	CreateResult
// @Tags Result
// @Description Create Result
// @ID create-result
// @Accept json
// @Produce	json
// @Param input	body request.Result true "body info"
// @Success 200 {object} response.Result
// @Router /api/v1/result/field/:fieldId [post]
func CreateResult(ctx *fiber.Ctx) error {
	fieldID, _ := ctx.ParamsInt("fieldId")

	body := new(request.Result)
	body.FieldID = uint(fieldID)

	if err := ctx.BodyParser(body); err != nil {
		log.Println("Error in parsing request", err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in parsing request",
		})
	}

	if res, err := service.CreateResult(*body); err != nil {
		log.Println("Error in result service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in result service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
}

// @Summary	CreateResults
// @Tags Result
// @Description Create Results
// @ID create-results
// @Accept json
// @Produce	json
// @Param input	body request.Result true "body info"
// @Success 200 {array} response.Result
// @Router /api/v1/result [post]
func CreateResults(ctx *fiber.Ctx) error {
	body := new([]request.Result)

	if err := ctx.BodyParser(body); err != nil {
		log.Println("Error in parsing request", err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in parsing request",
		})
	}

	if len(*body) < 1 {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Clear body",
		})
	}

	var res []response.Result

	for _, item := range *body {
		if result, err := service.CreateResult(item); err != nil {
			log.Println("Error in result service", err)
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Error in result service",
			})
		} else {
			res = append(res, result)
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": res,
	})
}

// @Summary	GetAllResults
// @Tags Result
// @Description Get All Results
// @ID get-all-results
// @Accept json
// @Produce	json
// @Success 200 {array} response.Result
// @Router /api/v1/result [get]
func GetAllResults(ctx *fiber.Ctx) error {
	if res, err := service.GetAllResults(); err != nil {
		log.Println("Error in result service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in result service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
}

// @Summary	GetOneResults
// @Tags Result
// @Description Get One Results
// @ID get-one-results
// @Accept json
// @Produce	json
// @Success 200 {object} response.Result
// @Router /api/v1/result/:id [get]
func GetOneResult(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	if res, err := service.GetOneResult(id); err != nil {
		log.Println("Error in result service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in result service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
}

// @Summary	GetFormResults
// @Tags Result
// @Description Get Form Results
// @ID get-form-results
// @Accept json
// @Produce	json
// @Success 200 {array} response.Result
// @Router /api/v1/result/form/:formId [get]
func GetFormResults(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("formId")

	if res, err := service.GetFormResults(id); err != nil {
		log.Println("Error in result service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in result service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
}

// @Summary	GetFieldResults
// @Tags Result
// @Description Get Field Results
// @ID get-field-results
// @Accept json
// @Produce	json
// @Success 200 {array} response.Result
// @Router /api/v1/result/field/:fieldId [get]
func GetFieldResults(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("fieldId")

	if res, err := service.GetFieldResults(id); err != nil {
		log.Println("Error in result service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in result service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
}

//// @Summary	UpdateResult
//// @Tags Permission
//// @Description Update Result
//// @ID update-result
//// @Accept json
//// @Produce	json
//// @Param input	body request.Result true "body info"
//// @Success 200 {object} response.Result
//// @Router /api/v1/result/:id [patch]
//func UpdateResult(ctx *fiber.Ctx) error {
//	body := new(request.ResultUpdate)
//	id, _ := ctx.ParamsInt("id")
//
//	if err := ctx.BodyParser(body); err != nil {
//		log.Println("Error in parsing request", err)
//		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
//			"message": "Error in parsing request",
//		})
//	}
//
//	if permission, err := service.UpdateResult(*body, id); err != nil {
//		log.Println("Error in result service", err)
//		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//			"message": "Error in result service",
//		})
//	} else {
//		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
//			"data": permission,
//		})
//	}
//}

// @Summary DeleteResult
// @Tags Result
// @Description Delete Result
// @ID delete-result
// @Accept json
// @Produce	json
// @Success 200 {string} string
// @Router /api/v1/result/:id [delete]
func DeleteResult(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	if err := service.DeleteResult(id); err != nil {
		log.Println("Error in result service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in result service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Result deleted",
		})
	}
}

// @Summary DeleteFormResult
// @Tags Result
// @Description Delete Form Result
// @ID delete-form-result
// @Accept json
// @Produce	json
// @Success 200 {string} string
// @Router /api/v1/result/form/:formId [delete]
func DeleteFormResults(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("formId")

	if err := service.DeleteFormResults(id); err != nil {
		log.Println("Error in result service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in result service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Result deleted",
		})
	}
}

// @Summary DeleteFieldResult
// @Tags Result
// @Description Delete Field Result
// @ID delete-field-result
// @Accept json
// @Produce	json
// @Success 200 {string} string
// @Router /api/v1/result/field/:fieldId [delete]
func DeleteFieldResults(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("fieldId")

	if err := service.DeleteFieldResults(id); err != nil {
		log.Println("Error in result service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in result service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Result deleted",
		})
	}
}
