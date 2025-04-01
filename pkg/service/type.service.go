package service

import (
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/request"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/response"
)

func CreateString(body map[string]interface{}, id int) (response.FieldString, error) {
	res := new(response.FieldString)

	if placeholder, err := CreateFieldPlaceholder(request.FieldPlaceholder{
		Placeholder: body["placeholder"].(string),
	}, id); err != nil {
		return response.FieldString{}, err
	} else {
		res.Placeholder = placeholder.Placeholder
	}

	if ranges, err := CreateFieldRange(request.FieldRange{
		Min: int(body["min"].(float64)),
		Max: int(body["max"].(float64)),
	}, id); err != nil {
		return response.FieldString{}, err
	} else {
		res.Min = ranges.Min
		res.Max = ranges.Max
	}

	return *res, nil
}

func GetString(id int) (response.FieldString, error) {
	res := new(response.FieldString)

	if placeholder, err := GetFieldPlaceholder(id); err != nil {
		return response.FieldString{}, err
	} else {
		res.Placeholder = placeholder.Placeholder
	}

	if ranges, err := GetFieldRange(id); err != nil {
		res.Min = ranges.Min
		res.Max = ranges.Max
	}

	return *res, nil
}

func CreateNumber(body map[string]interface{}, id int) (response.FieldNumber, error) {
	res := new(response.FieldNumber)

	if placeholder, err := CreateFieldPlaceholder(request.FieldPlaceholder{
		Placeholder: body["placeholder"].(string),
	}, id); err != nil {
		return response.FieldNumber{}, err
	} else {
		res.Placeholder = placeholder.Placeholder
	}

	if ranges, err := CreateFieldRange(request.FieldRange{
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

	if placeholder, err := GetFieldPlaceholder(id); err != nil {
		return response.FieldNumber{}, err
	} else {
		res.Placeholder = placeholder.Placeholder
	}

	if ranges, err := GetFieldRange(id); err != nil {
		res.Min = ranges.Min
		res.Max = ranges.Max
	}

	return *res, nil
}

func CreateEmail(body map[string]interface{}, id int) (response.FieldEmail, error) {
	res := new(response.FieldEmail)

	if placeholder, err := CreateFieldPlaceholder(request.FieldPlaceholder{
		Placeholder: body["placeholder"].(string),
	}, id); err != nil {
		return response.FieldEmail{}, err
	} else {
		res.Placeholder = placeholder.Placeholder
	}

	if ranges, err := CreateFieldRange(request.FieldRange{
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

	if placeholder, err := GetFieldPlaceholder(id); err != nil {
		return response.FieldEmail{}, err
	} else {
		res.Placeholder = placeholder.Placeholder
	}

	if ranges, err := GetFieldRange(id); err != nil {
		res.Min = ranges.Min
		res.Max = ranges.Max
	}

	return *res, nil
}

func CreateText(body map[string]interface{}, id int) (response.FieldText, error) {
	res := new(response.FieldText)

	if placeholder, err := CreateFieldPlaceholder(request.FieldPlaceholder{
		Placeholder: body["placeholder"].(string),
	}, id); err != nil {
		return response.FieldText{}, err
	} else {
		res.Placeholder = placeholder.Placeholder
	}

	if ranges, err := CreateFieldRange(request.FieldRange{
		Min: int(body["min"].(float64)),
		Max: int(body["max"].(float64)),
	}, id); err != nil {
		return response.FieldText{}, err
	} else {
		res.Min = ranges.Min
		res.Max = ranges.Max
	}

	return *res, nil
}

func GetText(id int) (response.FieldText, error) {
	res := new(response.FieldText)

	if placeholder, err := GetFieldPlaceholder(id); err != nil {
		return response.FieldText{}, err
	} else {
		res.Placeholder = placeholder.Placeholder
	}

	if ranges, err := GetFieldRange(id); err != nil {
		res.Min = ranges.Min
		res.Max = ranges.Max
	}

	return *res, nil
}

func CreateDate(body map[string]interface{}, id int) (response.FieldText, error) {
	res := new(response.FieldText)

	if placeholder, err := CreateFieldPlaceholder(request.FieldPlaceholder{
		Placeholder: body["placeholder"].(string),
	}, id); err != nil {
		return response.FieldText{}, err
	} else {
		res.Placeholder = placeholder.Placeholder
	}

	if ranges, err := CreateFieldRange(request.FieldRange{
		Min: int(body["min"].(float64)),
		Max: int(body["max"].(float64)),
	}, id); err != nil {
		return response.FieldText{}, err
	} else {
		res.Min = ranges.Min
		res.Max = ranges.Max
	}

	return *res, nil
}

func GetDate(id int) (response.FieldDate, error) {
	res := new(response.FieldDate)

	if placeholder, err := GetFieldPlaceholder(id); err != nil {
		return response.FieldDate{}, err
	} else {
		res.Placeholder = placeholder.Placeholder
	}

	if ranges, err := GetFieldRange(id); err != nil {
		res.Min = ranges.Min
		res.Max = ranges.Max
	}

	return *res, nil
}

func CreateRadio(body map[string]interface{}, id int) (response.FieldRadio, error) {
	res := new(response.FieldRadio)

	if multiply, err := CreateFieldMultiply(request.FieldMultiply{
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

	if multiply, err := GetFieldMultiply(id); err != nil {
		return response.FieldRadio{}, err
	} else {
		res.IsMultiply = multiply.IsMultiply
	}

	return *res, nil
}

func CreateSelect(body map[string]interface{}, id int) (response.FieldSelect, error) {
	res := new(response.FieldSelect)

	if placeholder, err := CreateFieldPlaceholder(request.FieldPlaceholder{
		Placeholder: body["placeholder"].(string),
	}, id); err != nil {
		return response.FieldSelect{}, err
	} else {
		res.Placeholder = placeholder.Placeholder
	}

	if multiply, err := CreateFieldMultiply(request.FieldMultiply{
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

	if placeholder, err := GetFieldPlaceholder(id); err != nil {
		return response.FieldSelect{}, err
	} else {
		res.Placeholder = placeholder.Placeholder
	}

	if multiply, err := GetFieldMultiply(id); err != nil {
		return response.FieldSelect{}, err
	} else {
		res.IsMultiply = multiply.IsMultiply
	}

	return *res, nil
}
