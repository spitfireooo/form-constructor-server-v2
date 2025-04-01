package response

type User struct {
	ID        uint    `json:"id" db:"id"`
	Email     string  `json:"email" db:"email"`
	Phone     *string `json:"phone" db:"phone"`
	Address   *string `json:"address" db:"address"`
	Nickname  *string `json:"nickname" db:"nickname"`
	Logo      *string `json:"logo" db:"logo"`
	CreatedAt string  `json:"created_at" db:"created_at"`
	UpdatedAt string  `json:"updated_at" db:"updated_at"`
}

type Token struct {
	ID     uint   `json:"id" db:"id"`
	UserID uint   `json:"user_id" db:"user_id"`
	Token  string `json:"token" db:"token"`
}

type UserPermission struct {
	ID         uint   `json:"id" db:"id"`
	UserID     uint   `json:"user_id" db:"user_id"`
	Permission string `json:"permission" db:"permission"`
}

type UserTag struct {
	ID     uint `json:"id" db:"id"`
	UserID uint `json:"user_id" db:"user_id"`
	TagID  uint `json:"tag_id" db:"tag_id"`
}

type Tag struct {
	ID          uint    `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description *string `json:"description" db:"description"`
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

type FieldString struct {
	Placeholder string `json:"placeholder" db:"placeholder"`
	Min         int    `json:"min" db:"min"`
	Max         int    `json:"max" db:"max"`
}

type FieldNumber struct {
	Placeholder string `json:"placeholder" db:"placeholder"`
	Min         int    `json:"min" db:"min"`
	Max         int    `json:"max" db:"max"`
}

type FieldEmail struct {
	Placeholder string `json:"placeholder" db:"placeholder"`
	Min         int    `json:"min" db:"min"`
	Max         int    `json:"max" db:"max"`
}

type FieldText struct {
	Placeholder string `json:"placeholder" db:"placeholder"`
	Min         int    `json:"min" db:"min"`
	Max         int    `json:"max" db:"max"`
}

type FieldDate struct {
	Placeholder string `json:"placeholder" db:"placeholder"`
	Min         int    `json:"min" db:"min"`
	Max         int    `json:"max" db:"max"`
}

type FieldRadio struct {
	IsMultiply bool `json:"is_multiply" db:"is_multiply"`
}

type FieldSelect struct {
	Placeholder string `json:"placeholder" db:"placeholder"`
	IsMultiply  bool   `json:"is_multiply" db:"is_multiply"`
}
