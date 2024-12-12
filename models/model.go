package models

type CreateUserPayload struct {
	UserName       string `json:"user_name"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

type CreateGroupPayload struct {
	GroupName string `json:"group_name"`
	CreatorID int    `json:"created_by"`
}
