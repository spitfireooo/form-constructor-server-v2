package service_v2

import (
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/request"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/response"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/service"
)

func CreateNumber(body map[string]interface{}, id int) (response.FieldNumber, error) {
	res := new(response.FieldNumber)

	if placeholder, err := service.CreateFieldPlaceholder(request.FieldPlaceholder{
		Placeholder: body["placeholder"].(string),
	}, id); err != nil {
		return response.FieldNumber{}, err
	} else {
		res.Placeholder = placeholder.Placeholder
	}

	if ranges, err := service.CreateFieldRange(request.FieldRange{
		Min: int(body["min"].(float64)),
		Max: int(body["max"].(float64)),
	}, id); err != nil {
		return response.FieldNumber{}, err
	} else {
		res.Min = ranges.Min
		res.Max = ranges.Max
	}

	return *res, nil
}

func GetNumber(id int) (response.FieldNumber, error) {
	res := new(response.FieldNumber)

	if placeholder, err := service.GetFieldPlaceholder(id); err != nil {
		return response.FieldNumber{}, err
	} else {
		res.Placeholder = placeholder.Placeholder
	}

	if ranges, err := service.GetFieldRange(id); err != nil {
		return response.FieldNumber{}, err
	} else {
		res.Min = ranges.Min
		res.Max = ranges.Max
	}

	return *res, nil
}

func UpdateNumber(body map[string]interface{}, id int) (response.FieldNumber, error) {
	res := new(response.FieldNumber)

	if placeholder, err := service.UpdateFieldPlaceholder(request.FieldPlaceholder{
		Placeholder: body["placeholder"].(string),
	}, id); err != nil {
		return response.FieldNumber{}, err
	} else {
		res.Placeholder = placeholder.Placeholder
	}

	Min := int(body["min"].(float64))
	Max := int(body["max"].(float64))
	if ranges, err := service.UpdateFieldRange(request.FieldRangeUpdate{
		Min: &Min,
		Max: &Max,
	}, id); err != nil {
		return response.FieldNumber{}, err
	} else {
		res.Min = ranges.Min
		res.Max = ranges.Max
	}

	return *res, nil
}
