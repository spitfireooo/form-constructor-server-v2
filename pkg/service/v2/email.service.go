package service_v2

import (
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/request"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/response"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/service"
)

func CreateEmail(body map[string]interface{}, id int) (response.FieldEmail, error) {
	res := new(response.FieldEmail)

	if placeholder, err := service.CreateFieldPlaceholder(request.FieldPlaceholder{
		Placeholder: body["placeholder"].(string),
	}, id); err != nil {
		return response.FieldEmail{}, err
	} else {
		res.Placeholder = placeholder.Placeholder
	}

	if ranges, err := service.CreateFieldRange(request.FieldRange{
		Min: int(body["min"].(float64)),
		Max: int(body["max"].(float64)),
	}, id); err != nil {
		return response.FieldEmail{}, err
	} else {
		res.Min = ranges.Min
		res.Max = ranges.Max
	}

	return *res, nil
}

func GetEmail(id int) (response.FieldEmail, error) {
	res := new(response.FieldEmail)

	if placeholder, err := service.GetFieldPlaceholder(id); err != nil {
		return response.FieldEmail{}, err
	} else {
		res.Placeholder = placeholder.Placeholder
	}

	if ranges, err := service.GetFieldRange(id); err != nil {
		return response.FieldEmail{}, err
	} else {
		res.Min = ranges.Min
		res.Max = ranges.Max
	}

	return *res, nil
}

func UpdateEmail(body map[string]interface{}, id int) (response.FieldEmail, error) {
	res := new(response.FieldEmail)

	if placeholder, err := service.UpdateFieldPlaceholder(request.FieldPlaceholder{
		Placeholder: body["placeholder"].(string),
	}, id); err != nil {
		return response.FieldEmail{}, err
	} else {
		res.Placeholder = placeholder.Placeholder
	}

	if ranges, err := service.UpdateFieldRange(request.FieldRangeUpdate{
		Min: body["min"].(*int),
		Max: body["max"].(*int),
	}, id); err != nil {
		return response.FieldEmail{}, err
	} else {
		res.Min = ranges.Min
		res.Max = ranges.Max
	}

	return *res, nil
}
