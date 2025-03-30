package request

type Result struct {
	Value   string `json:"value" db:"value"`
	FieldID uint   `json:"field_id" db:"field_id"`
	UserID  uint   `json:"user_id" db:"user_id"`
}

type ResultUpdate struct {
	Value   *string `json:"value,omitempty" db:"value"`
	FieldID *uint   `json:"field_id,omitempty" db:"field_id"`
	UserID  *uint   `json:"user_id,omitempty" db:"user_id"`
}
