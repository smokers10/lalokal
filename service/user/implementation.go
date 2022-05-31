package user

import (
	"fmt"
	"lalokal/domain/forgot_password"
	"lalokal/domain/http_response"
	"lalokal/domain/user"
	"lalokal/domain/verification"
	"lalokal/infrastructure/configuration"
	"lalokal/infrastructure/dependency/encryption"
	"lalokal/infrastructure/dependency/identifier"
	"lalokal/infrastructure/dependency/jsonwebtoken"
	"lalokal/infrastructure/lib"
	"lalokal/infrastructure/lib/smtp"
	"lalokal/service/user/helper"
)

type userService struct {
	userRepository           user.Repository
	forgotPasswordRepository forgot_password.Repository
	verificationRepository   verification.Repository
	bcrypt                   encryption.Contract
	jsonwebtoken             jsonwebtoken.Contact
	identifier               identifier.Contract
	smtp                     smtp.Contract
}

// ForgotPassword implements user.Service
func (s *userService) ForgotPassword(email string) (response *http_response.Response) {
	if msg, isfail := emailValidation(email); isfail {
		return &http_response.Response{Message: msg, Status: 400}
	}

	// retrieve user by email
	user := s.userRepository.FindOneByEmail(email)

	// when user not exists/registered
	if user.Id == "" {
		return &http_response.Response{
			Message: "pengguna tidak terdaftar",
			Status:  404,
		}
	}

	// create reset password secret code/OTP
	otp := lib.OTPGenerator()
	token := s.identifier.MakeIdentifier()
	hashed_otp := s.bcrypt.Hash(otp)

	// store forgot password data
	if err := s.forgotPasswordRepository.Insert(&forgot_password.ForgotPassword{
		Token:  token,
		UserId: user.Id,
		Secret: hashed_otp,
	}); err != nil {
		return &http_response.Response{
			Message: "kesalahan saat menyimpan data atur ulang password",
			Status:  500,
		}
	}

	// make link
	configuration := configuration.ReadConfiguration().Application
	link := fmt.Sprintf("%s/%s", configuration.BaseURL, token)

	// send forgot password email
	if err := s.smtp.Send([]string{user.Email}, "Atur Ulang Password Lalokal", helper.ForgotPasswordEmailTemplate(user.Name, otp, link)); err != nil {
		return &http_response.Response{
			Message: "kesalahan saat mengirim email atur ulang password",
			Status:  500,
		}
	}

	return &http_response.Response{
		Message: "link atur password telah dikirim",
		Success: true,
		Status:  200,
	}
}

// GetProfile implements user.Service
func (s *userService) GetProfile(user_id string) (response *http_response.Response) {
	if user_id == "" {
		return &http_response.Response{
			Message: "id pengguna tidak boleh kosong",
			Status:  400,
		}
	}

	result := s.userRepository.FindOneById(user_id)

	return &http_response.Response{
		Message: "profile berhasil diambil",
		Success: true,
		Status:  200,
		Data:    result,
	}
}

// Login implements user.Service
func (s *userService) Login(input *user.LoginData) (response *http_response.Response) {
	if msg, isfail := loginValidation(input); isfail {
		return &http_response.Response{Message: msg, Status: 400}
	}

	// retrieve user
	user := s.userRepository.FindOneByEmail(input.Email)
	if user.Id == "" {
		return &http_response.Response{
			Message: "email atau password salah",
			Status:  401,
		}
	}

	// compare password
	if !s.bcrypt.Compare(user.Password, input.Password) {
		return &http_response.Response{
			Message: "email atau password salah",
			Status:  401,
		}
	}

	// make token
	payload := map[string]interface{}{
		"id":    user.Id,
		"email": user.Email,
	}
	token, err := s.jsonwebtoken.Sign(payload)
	if err != nil {
		return &http_response.Response{
			Message: "kesalahan saat membuat token",
			Status:  500,
		}
	}

	// assign empty string to password in order to hide it
	user.Password = ""

	return &http_response.Response{
		Message: "login berhasil",
		Success: true,
		Status:  200,
		Token:   token,
		Data:    user,
	}
}

// Register implements user.Service
func (s *userService) Register(input *user.RegisterData) (response *http_response.Response) {
	if msg, isfail := registerValidation(input); isfail {
		return &http_response.Response{Message: msg, Status: 400}
	}

	// retrieve verification data
	verification := s.verificationRepository.FindOneByEmail(input.Email)

	// if email not verified
	if verification.Status == "not verified" || verification.Status == "" {
		return &http_response.Response{
			Message: "email tidak tervalidasi",
			Status:  401,
		}
	}

	// retrieve vrification data
	user := s.userRepository.FindOneByEmail(input.Email)

	// if user already registered
	if user.Id != "" {
		return &http_response.Response{
			Message: "pengguna sudah terdaftar",
			Status:  401,
		}
	}

	// secure password
	input.Password = s.bcrypt.Hash(input.Password)

	// store user
	inserted_id, err := s.userRepository.Insert(input)
	if err != nil {
		return &http_response.Response{
			Message: "kesalahan saat menyimpan pengguna",
			Status:  500,
		}
	}

	// make new token
	payload := map[string]interface{}{
		"id":    inserted_id,
		"email": input.Email,
	}
	token, err := s.jsonwebtoken.Sign(payload)
	if err != nil {
		return &http_response.Response{
			Message: "kesalahan saat membuat token",
			Status:  500,
		}
	}

	// safe response
	input.Password = ""
	input.CofirmPassword = ""

	return &http_response.Response{
		Message: "registrasi berhasil",
		Success: true,
		Status:  200,
		Token:   token,
		Data:    input,
	}
}

// ResetPassword implements user.Service
func (s *userService) ResetPassword(input *user.ResetPasswordData) (response *http_response.Response) {
	if msg, isfail := resetPasswordValidation(input); isfail {
		return &http_response.Response{Message: msg, Status: 400}
	}

	panic("unimplemented")
}

// UpdateProfile implements user.Service
func (s *userService) UpdateProfile(input *user.User) (response *http_response.Response) {
	if msg, isfail := updateProfileValidation(input); isfail {
		return &http_response.Response{Message: msg, Status: 400}
	}

	panic("unimplemented")
}

// VerificateEmail implements user.Service
func (s *userService) VerificateEmail(input *verification.Verification) (response *http_response.Response) {
	if msg, isfail := vefiricateEmailValidation(input); isfail {
		return &http_response.Response{Message: msg, Status: 400}
	}

	panic("unimplemented")
}

// VerificationRequest implements user.Service
func (s *userService) VerificationRequest(email string) (response *http_response.Response) {
	if msg, isfail := emailValidation(email); isfail {
		return &http_response.Response{Message: msg, Status: 400}
	}

	panic("unimplemented")
}

func UserService(ur *user.Repository, fp *forgot_password.Repository, vr *verification.Repository,
	en *encryption.Contract, jwt *jsonwebtoken.Contact, id *identifier.Contract, smtp *smtp.Contract) user.Service {
	return &userService{
		userRepository:           *ur,
		forgotPasswordRepository: *fp,
		verificationRepository:   *vr,
		bcrypt:                   *en,
		jsonwebtoken:             *jwt,
		identifier:               *id,
		smtp:                     *smtp,
	}
}
