package idcloudhost

type ProfileInput struct {
	Firstname        string `form:"first_name"`
	Lastname         string `form:"last_name"`
	PhoneNumber      string `form:"phone_number"`
	PersonalIDNumber string `form:"personal_id_number"`
}
