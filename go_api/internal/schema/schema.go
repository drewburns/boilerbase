package schema

// Schema example:

// type User struct {
// 	ID           int    `json:"id" pg:"id, pk"`
// 	Username     string `json:"username" pg:"username"`
// 	Email        string `pg:"email" json:"-"`
// 	PasswordHash string `json:"-" pg:"encrypted_password"`
// 	// Posts                *[]Post
// 	CreatedAt     time.Time `pg:"created_at"`
// 	UpdatedAt     time.Time `pg:"updated_at"`
// }

// type Post struct {
// 	ID        int    `json:"id" pg:"id, pk"`
// 	Caption   string `pg:"caption" json:"caption"`
// 	UserID    int    `pg:"user_id, fk" json:"user_id"`
// 	User      *User
// 	CreatedAt time.Time `pg:"created_at" json:"created_at"`
// 	UpdatedAt time.Time `pg:"updated_at"`
// }
