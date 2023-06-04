package models

//InfoBody /info body
type InfoBody struct {
	Username string
}

//InstagramUser instagram user
type InstagramUser struct {
	ID             int64  `json:"id"`
	Username       string `json:"username"`
	FullName       string `json:"full_name"`
	Bio            string `json:"bio"`
	ProfilePicture string `json:"profile_picture"`
	Private        bool   `json:"private"`
	Followers      int    `json:"followers"`
}
