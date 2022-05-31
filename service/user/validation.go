package user

import (
	"lalokal/domain/user"
	"lalokal/domain/verification"
	"net"
	"net/mail"
	"strings"
)

func emailValidation(email string) (message string, isfail bool) {
	if email == "" {
		return "email tidak boleh kosong", true
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return "email tidak valid", true
	}

	domainPart := strings.Split(email, "@")
	_, err = net.LookupMX(domainPart[1])
	if err != nil {
		return "domain email tidak terdaftar", true
	}

	return "", false
}

func loginValidation(i *user.LoginData) (message string, isfail bool) {
	if msg, isfail := emailValidation(i.Email); isfail {
		return msg, isfail
	}

	if i.Password == "" {
		return "password tidak boleh kosong", true
	}

	return "", false
}

func registerValidation(i *user.RegisterData) (message string, isfail bool) {
	if msg, isfail := emailValidation(i.Email); isfail {
		return msg, isfail
	}

	if i.Name == "" {
		return "nama tidak boleh kosong", true
	}

	if i.CompanyName == "" {
		return "nama perusahaan tidak boleh kosong", true
	}

	if i.Password == "" {
		return "password tidak boleh kosong", true
	}

	if i.CofirmPassword == "" {
		return "konfirmasi password tidak boleh kosong", true
	}

	if len(i.Password) < 8 || len(i.CofirmPassword) < 0 {
		return "panjang password harus lebih dari 8 karakter", true
	}

	if i.Password != i.CofirmPassword {
		return "konfirmasi password salah", true
	}

	return "", false
}

func resetPasswordValidation(i *user.ResetPasswordData) (message string, isfail bool) {
	if i.Password == "" {
		return "password tidak boleh kosong", true
	}

	if i.CofirmPassword == "" {
		return "konfirmasi password tidak boleh kosong", true
	}

	if len(i.Password) < 8 || len(i.CofirmPassword) < 0 {
		return "panjang password harus lebih dari 8 karakter", true
	}

	if i.Password != i.CofirmPassword {
		return "konfirmasi password salah", true
	}

	if i.Secret == "" {
		return "kode rahasia tidak boleh kosong", true
	}

	if i.Token == "" {
		return "token reset password tidak boleh kosong", true
	}

	return "", false
}

func updateProfileValidation(i *user.User) (message string, isfail bool) {
	if i.Name == "" {
		return "nama tidak boleh kosong", true
	}

	if i.CompanyName == "" {
		return "nama perusahaan tidak boleh kosong", true
	}

	if i.Id == "" {
		return "id tidak boleh kosong", true
	}

	return "", false
}

func vefiricateEmailValidation(i *verification.Verification) (message string, isfail bool) {
	if i.RequesterEmail == "" {
		return "email pengaju tidak boleh kosong", true
	}

	if i.Secret == "" {
		return "kode OTP tidak boleh kosong", true
	}

	return "", false
}
