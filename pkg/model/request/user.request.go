package request

type User struct {
	Email    string  `json:"email" db:"email"`
	Phone    *string `json:"phone" db:"phone"`
	Address  *string `json:"address" db:"address"`
	Password *string `json:"password" db:"password"`
	Nickname *string `json:"nickname" db:"nickname"`
	Logo     string  `json:"logo" db:"logo"`
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
