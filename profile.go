package idcloudhost

type Profile struct {
	// Avatar Image URL
	Avatar string `json:"avatar"`

	// User created date
	CreatedAt string `json:"created_at"`

	// User email
	Email string `json:"email"`

	// User first name
	FirstName string `json:"first_name"`

	// Profile ID
	ID int64 `json:"id"`

	// User last name
	LastName string `json:"last_name"`

	// Personal identity number, usually issued by the state
	PersonalIDNumber string `json:"personal_id_number"`

	// Phone number in any format
	PhoneNumber string `json:"phone_number"`

	// User updated date
	UpdatedAt string `json:"updated_at"`

	// User ID
	UserID int64 `json:"user_id"`
}
