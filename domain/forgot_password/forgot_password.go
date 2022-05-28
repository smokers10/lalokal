package forgot_password

type ForgotPassword struct {
	Id     string `json:"id,omitempty" form:"id" bson:"_id"`
	Token  string `json:"token,omitempty" form:"token" bson:"token"`
	UserId string `json:"user_id,omitempty" form:"user_id" bson:"user_id"`
	Secret string `json:"secret,omitempty" form:"secret" bson:"secret"`
}

type Repository interface {
	Insert(data *ForgotPassword) (failure error)

	FindOneByToken(token string) (result *ForgotPassword)

	Delete(token string) (failure error)
}
