package response

type Result struct {
	ID      uint   `json:"id" db:"id"`
	Type    string `json:"type" db:"type"`
	Value   any    `json:"value" db:"value"`
	FormID  uint   `json:"form_id" db:"form_id"`
	FieldID uint   `json:"field_id" db:"field_id"`
	UserID  uint   `json:"user_id" db:"user_id"`
}
