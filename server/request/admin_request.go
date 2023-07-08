package request

type Admin_CreateUser struct {
	Name string `json:"name" binding:"required"`
}
