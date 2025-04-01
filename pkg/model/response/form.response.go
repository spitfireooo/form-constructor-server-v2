package response

type Form struct {
	ID          uint    `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Slug        string  `json:"slug" db:"slug"`
	Description *string `json:"description" db:"description"`
	Logo        *string `json:"logo" db:"logo"`
	AuthorId    uint    `json:"author_id" db:"author_id"`
}

type Field struct {
	ID       uint   `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Type     string `json:"type" db:"type"`
	Label    string `json:"label" db:"label"`
	OrderOf  int    `json:"order_of" db:"order_of"`
	Required bool   `json:"required" db:"required"`
	FormID   uint   `json:"form_id" db:"form_id"`
}
