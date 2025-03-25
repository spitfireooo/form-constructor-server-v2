package request

type User struct {
	Email    string  `json:"email,omitempty" db:"email" validate:"omitempty,email"`
	Phone    *string `json:"phone,omitempty" db:"phone" validate:"omitempty,min=3,max=20"`
	Address  *string `json:"address,omitempty" db:"address" validate:"omitempty,min=3"`
	Password *string `json:"password,omitempty" db:"password" validate:"omitempty,min=6,max=30"`
	Nickname *string `json:"nickname,omitempty" db:"nickname" validate:"omitempty,min=3,max=30"`
	Logo     string  `json:"logo,omitempty" db:"logo" validate:"omitempty,image"`
}

type Token struct {
	Token string `json:"token" db:"token"`
}

type UserPermission struct {
	Permission string `json:"permission" db:"permission"`
}

type UserTag struct {
	TagID uint `json:"tag_id" db:"tag_id"`
}

type Tag struct {
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}
