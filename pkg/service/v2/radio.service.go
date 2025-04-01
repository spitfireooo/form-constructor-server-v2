package service_v2

import (
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/request"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/response"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/service"
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

	return *res, nil
}

func GetRadio(id int) (response.FieldRadio, error) {
	res := new(response.FieldRadio)

	if multiply, err := service.GetFieldMultiply(id); err != nil {
		return response.FieldRadio{}, err
	} else {
		res.IsMultiply = multiply.IsMultiply
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

	return *res, nil
}
