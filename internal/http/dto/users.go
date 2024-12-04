package dto

type GetUserByIDRequest struct {
	ID int64 `param:"id" validate:"required"`
}
type UserLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`

}

type UserRegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	FullName string `json:"full_name" validate:"required"`
}
