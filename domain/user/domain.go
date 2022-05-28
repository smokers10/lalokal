package user

type User struct {
	Id          string `json:"id,omitempty" bson:"_id"`
	Name        string `json:"name,omitempty" bson:"name"`
	CompanyName string `json:"company_name,omitempty" bson:"company_name"`
	Email       string `json:"email,omitempty" bson:"email"`
	Password    string `json:"password,omitempty" bson:"password"`
}

type RegisterData struct {
	Name           string `json:"name" form:"name"`
	CompanyName    string `json:"company_name" form:"company_name"`
	Email          string `json:"email" form:"email"`
	Password       string `json:"password" form:"password"`
	CofirmPassword string `json:"confirm_password" form:"confirm_password"`
}

type VerificateEmailData struct {
	Email  string `json:"email" form:"email"`
	Secret string `json:"secret" form:"secret"`
}

type ResetPasswordData struct {
	Token          string `json:"token" form:"token"`
	Secret         string `json:"secret" form:"secret"`
	Password       string `json:"password" form:"password"`
	CofirmPassword string `json:"confirm_password" form:"confirm_password"`
	UserId         string `json:"user_id" form:"user_id"`
}

type LoginData struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
