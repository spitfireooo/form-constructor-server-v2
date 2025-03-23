package request

type Form struct {
	Title       string `json:"title" db:"title"`
	Slug        string `json:"slug" db:"slug"`
	Description uint   `json:"description" db:"description"`
	Logo        string `json:"logo" db:"logo"`
}

type Field struct {
	Name     string `json:"name" db:"name"`
	Type     string `json:"type" db:"type"`
	Label    string `json:"label" db:"label"`
	OrderOf  int    `json:"order_of" db:"order_of"`
	Required bool   `json:"required" db:"required"`
}

type FieldVariants struct {
	Variant string `json:"variant" db:"variant"`
	Name    string `json:"name" db:"name"`
}

type FieldRange struct {
	Min int `json:"min" db:"min"`
	Max int `json:"max" db:"max"`
}

type FieldPlaceholder struct {
	Placeholder string `json:"placeholder" db:"placeholder"`
}

type FieldMultiply struct {
	IsMultiply bool `json:"is_multiply" db:"is_multiply"`
}
