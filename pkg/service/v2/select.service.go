package service_v2

import (
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/request"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/response"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/service"
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
