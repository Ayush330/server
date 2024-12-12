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

type AddToGroupPayload struct {
	GroupId int `json:"group_id"`
	UserId  int `json:"user_id"`
}

type AddExpensePayload struct {
	GroupId         int                   `json:"group_id"`
	PaidBy          int                   `json:"paid_by"`
	Description     string                `json:"description"`
	TotalAmount     int                   `json:"total_amount"`
	UserListToSplit []ExpenseSplitPayload `json:"user_list_to_split"`
}

type ExpenseSplitPayload struct {
	UserId int `json:"user_id"`
	Amount int `json:"amount"`
}
