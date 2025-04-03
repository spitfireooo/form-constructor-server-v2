package request

type FormWithField struct {
	Title       string  `json:"title" db:"title" validate:"min=3,max=30"`
	Slug        string  `json:"slug" db:"slug" validate:"min=3,max=50"`
	Description *string `json:"description,omitempty" db:"description" validate:"omitempty,min=5"`
	Logo        *string `json:"logo,omitempty" db:"logo" validate:"omitempty,image"`

	Fields *[]map[string]interface{} `json:"fields,omitempty" db:"fields" validate:"omitempty,fields"`
}

type Form struct {
	Title       string  `json:"title" db:"title" validate:"min=3,max=30"`
	Slug        string  `json:"slug" db:"slug" validate:"min=3,max=50"`
	Description *string `json:"description,omitempty" db:"description" validate:"omitempty,min=5"`
	Logo        *string `json:"logo,omitempty" db:"logo" validate:"omitempty,image"`
}

type FormUpdate struct {
	Title       *string `json:"title,omitempty" db:"title" validate:"omitempty,min=3,max=30"`
	Slug        *string `json:"slug,omitempty" db:"slug" validate:"omitempty,min=3,max=50"`
	Description *string `json:"description,omitempty" db:"description" validate:"omitempty,min=5"`
	Logo        *string `json:"logo,omitempty" db:"logo" validate:"omitempty,image"`
	AuthorId    *uint   `json:"author_id,omitempty" db:"logo" validate:"omitempty"`
}

type Field struct {
	Name     string `json:"name" db:"name" validate:"min=2"`
	Type     string `json:"type" db:"type"`
	Label    string `json:"label" db:"label" validate:"min=2"`
	OrderOf  int    `json:"order_of" db:"order_of"`
	Required bool   `json:"required" db:"required"`
}

type FieldUpdate struct {
	FormID   *uint   `json:"form_id,omitempty" db:"form_id" validate:"omitempty"`
	Name     *string `json:"name,omitempty" db:"name" validate:"omitempty,min=2"`
	Type     *string `json:"type,omitempty" db:"type" validate:"omitempty"`
	Label    *string `json:"label,omitempty" db:"label" validate:"omitempty,min=2"`
	OrderOf  *int    `json:"order_of,omitempty" db:"order_of" validate:"omitempty"`
	Required *bool   `json:"required,omitempty" db:"required" validate:"omitempty"`
}

type FieldVariants struct {
	Variant string `json:"variant" db:"variant"`
	Name    string `json:"name" db:"name"`
}

type FieldVariantsUpdate struct {
	Variant *string `json:"variant,omitempty" db:"variant" validate:"omitempty,min=3"`
	Name    *string `json:"name,omitempty" db:"name" validate:"omitempty,min=3"`
}

type FieldRange struct {
	Min int `json:"min" db:"min"`
	Max int `json:"max" db:"max"`
}

type FieldRangeUpdate struct {
	Min *int `json:"min,omitempty" db:"min"`
	Max *int `json:"max,omitempty" db:"max"`
}

type FieldPlaceholder struct {
	Placeholder string `json:"placeholder" db:"placeholder"`
}

type FieldMultiply struct {
	IsMultiply bool `json:"is_multiply" db:"is_multiply"`
}

type FieldString struct {
	Name        string `json:"name" db:"name"`
	Type        string `json:"type" db:"type"`
	Label       string `json:"label" db:"label"`
	OrderOf     int    `json:"order_of" db:"order_of"`
	Required    bool   `json:"required" db:"required"`
	Placeholder string `json:"placeholder" db:"placeholder"`
	Min         int    `json:"min" db:"min"`
	Max         int    `json:"max" db:"max"`
}

type FieldNumber struct {
	Name        string `json:"name" db:"name"`
	Type        string `json:"type" db:"type"`
	Label       string `json:"label" db:"label"`
	OrderOf     int    `json:"order_of" db:"order_of"`
	Required    bool   `json:"required" db:"required"`
	Placeholder string `json:"placeholder" db:"placeholder"`
	Min         int    `json:"min" db:"min"`
	Max         int    `json:"max" db:"max"`
}

type FieldEmail struct {
	Name        string `json:"name" db:"name"`
	Type        string `json:"type" db:"type"`
	Label       string `json:"label" db:"label"`
	OrderOf     int    `json:"order_of" db:"order_of"`
	Required    bool   `json:"required" db:"required"`
	Placeholder string `json:"placeholder" db:"placeholder"`
	Min         int    `json:"min" db:"min"`
	Max         int    `json:"max" db:"max"`
}

type FieldText struct {
	Name        string `json:"name" db:"name"`
	Type        string `json:"type" db:"type"`
	Label       string `json:"label" db:"label"`
	OrderOf     int    `json:"order_of" db:"order_of"`
	Required    bool   `json:"required" db:"required"`
	Placeholder string `json:"placeholder" db:"placeholder"`
	Min         int    `json:"min" db:"min"`
	Max         int    `json:"max" db:"max"`
}

type FieldDate struct {
	Name        string `json:"name" db:"name"`
	Type        string `json:"type" db:"type"`
	Label       string `json:"label" db:"label"`
	OrderOf     int    `json:"order_of" db:"order_of"`
	Required    bool   `json:"required" db:"required"`
	Placeholder string `json:"placeholder" db:"placeholder"`
	Min         int    `json:"min" db:"min"`
	Max         int    `json:"max" db:"max"`
}

type FieldRadio struct {
	Name       string `json:"name" db:"name"`
	Type       string `json:"type" db:"type"`
	Label      string `json:"label" db:"label"`
	OrderOf    int    `json:"order_of" db:"order_of"`
	Required   bool   `json:"required" db:"required"`
	IsMultiply bool   `json:"is_multiply" db:"is_multiply"`
}

type FieldSelect struct {
	Name        string `json:"name" db:"name"`
	Type        string `json:"type" db:"type"`
	Label       string `json:"label" db:"label"`
	OrderOf     int    `json:"order_of" db:"order_of"`
	Required    bool   `json:"required" db:"required"`
	Placeholder string `json:"placeholder" db:"placeholder"`
	IsMultiply  bool   `json:"is_multiply" db:"is_multiply"`
}
