package dto

type CreateUserRequest struct {
	Username     string `json:"username" validate:"required,min=3,max=50,name_chars"`
	Email        string `json:"email" validate:"required,email"`
	Age          int    `json:"age" validate:"omitempty,gte=18"`
	PasswordHard string `json:"password" validate:"required,gte=7,lte=24,pass_chars"`
}

type UpdateUserRequest struct {
	Username string `json:"username" validate:"omitempty,min=3,max=50,name_chars"`
	Email    string `json:"email" validate:"omitempty,email"`
	Age      int    `json:"age" validate:"omitempty,gte=18"`
}

type GetUserReponse struct {
	Message string       `json:"msg"`
	Data    *UserReponse `json:"data"`
}

type GetAllUserReponse struct {
	Message string        `json:"msg"`
	Data    []UserReponse `json:"data"`
}

type UpdateUseReponse struct {
	Message string       `json:"msg"`
	Data    *UserReponse `json:"data"`
}
type DeleteReponse struct {
	Message string `json:"msg"`
	Dele    bool   `json:"delete"`
}
type CreateUserReponse struct {
	Message string       `json:"msg"`
	Data    *UserReponse `json:"data"`
}

type UserReponse struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	CreatedAt string `json:"Created_at"`
	UpdateAt  string `json:"Updated_at"`
}
