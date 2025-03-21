package entity

type Form struct {
	ID          uint   `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Slug        string `json:"slug" db:"slug"`
	Description uint   `json:"description" db:"description"`
	Logo        string `json:"logo" db:"logo"`
	AuthorId    uint   `json:"author_id" db:"author_id"`
}

type Field struct {
	ID       uint   `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Type     string `json:"type" db:"type"`
	Label    string `json:"label" db:"label"`
	OrderOf  int    `json:"order_of" db:"order_of"`
	Required bool   `json:"required" db:"required"`
	FormId   uint   `json:"form_id" db:"form_id"`
}

type FieldVariants struct {
	ID      uint   `json:"id" db:"id"`
	Variant string `json:"variant" db:"variant"`
	Name    string `json:"name" db:"name"`
	FieldId uint   `json:"field_id" db:"field_id"`
}

type FieldRange struct {
	ID      uint `json:"id" db:"id"`
	Min     int  `json:"min" db:"min"`
	Max     int  `json:"max" db:"max"`
	FieldId uint `json:"field_id" db:"field_id"`
}

type FieldPlaceholder struct {
	ID          uint   `json:"id" db:"id"`
	Placeholder string `json:"placeholder" db:"placeholder"`
	FieldId     uint   `json:"field_id" db:"field_id"`
}

type FieldMultiply struct {
	ID         uint `json:"id" db:"id"`
	IsMultiply bool `json:"is_multiply" db:"is_multiply"`
	FieldId    uint `json:"field_id" db:"field_id"`
}
