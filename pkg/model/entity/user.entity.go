package entity

type User struct {
	ID        uint    `json:"id" db:"id"`
	Email     string  `json:"email" db:"email"`
	Phone     *string `json:"phone" db:"phone"`
	Address   *string `json:"address" db:"address"`
	Password  string  `json:"password" db:"password"`
	Nickname  *string `json:"nickname" db:"nickname"`
	Logo      string  `json:"logo" db:"logo"`
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
	ID          uint   `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}
