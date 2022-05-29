package user

import "lalokal/domain/http_response"

type Repository interface {
	Insert(data *RegisterData) (failure error)

	UpdatePassword(data *ResetPasswordData) (failure error)

	Update(data *User) (failure error)

	FindOneByEmail(email string) (result *User)

	FindOneById(user_id string) (result *User)
}

type Service interface {
	Register(input *RegisterData) (response *http_response.Response)

	Login(input *LoginData) (response *http_response.Response)

	ForgotPassword(email string) (response *http_response.Response)

	ResetPassword(input *ResetPasswordData) (response *http_response.Response)

	GetProfile(user_id string) (response *http_response.Response)

	UpdateProfile(input *User) (response *http_response.Response)
}
