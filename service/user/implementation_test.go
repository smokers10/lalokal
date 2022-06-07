package user

import (
	"errors"
	"lalokal/domain/forgot_password"
	"lalokal/domain/user"
	"lalokal/domain/verification"
	"lalokal/infrastructure/encryption"
	"lalokal/infrastructure/identifier"
	"lalokal/infrastructure/jsonwebtoken"
	"lalokal/infrastructure/lib/common_testing"
	"lalokal/infrastructure/mailer"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userRepo           = user.MockRepository{Mock: mock.Mock{}}
	forgotPasswordRepo = forgot_password.MockRepository{Mock: mock.Mock{}}
	verificationRepo   = verification.MockRepository{Mock: mock.Mock{}}
	bcrypt             = encryption.MockContract{Mock: mock.Mock{}}
	jsonWebToken       = jsonwebtoken.MockContract{Mock: mock.Mock{}}
	id                 = identifier.MockContract{Mock: mock.Mock{}}
	smtpMailer         = mailer.MockContract{Mock: mock.Mock{}}
	service            = userService{
		userRepository:           &userRepo,
		forgotPasswordRepository: &forgotPasswordRepo,
		verificationRepository:   &verificationRepo,
		bcrypt:                   &bcrypt,
		jsonwebtoken:             &jsonWebToken,
		identifier:               &id,
		smtp:                     &smtpMailer,
	}
	emailTestTable = []struct {
		label    string
		email    string
		expected common_testing.Expectation
	}{
		{
			label: "empty email",
			email: "",
			expected: common_testing.Expectation{
				Message: "email tidak boleh kosong",
				Status:  400,
			},
		},
		{
			label: "invalid email",
			email: "johndoe@gmail@com",
			expected: common_testing.Expectation{
				Message: "email tidak valid",
				Status:  400,
			},
		},
		{
			label: "unregistered email domain",
			email: "johndoe@finalsacrophage.com",
			expected: common_testing.Expectation{
				Message: "domain email tidak terdaftar",
				Status:  400,
			},
		},
	}
)

func TestUserSerivce(t *testing.T) {
	s := UserService(&service.userRepository, &service.forgotPasswordRepository, &service.verificationRepository, &service.bcrypt, &service.jsonwebtoken, &service.identifier, &service.smtp)
	assert.NotEmpty(t, s)
}

func TestForgotPassword(t *testing.T) {
	t.Run("email testing", func(t *testing.T) {
		for _, tb := range emailTestTable {
			res := service.ForgotPassword(tb.email)

			common_testing.Assertion(t, tb.expected, res, common_testing.DefaultOption)
		}
	})

	t.Run("user not registerd", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "pengguna tidak terdaftar",
			Status:  404,
		}

		userRepo.Mock.On("FindOneByEmail", mock.Anything).Return(&user.User{}).Once()

		res := service.ForgotPassword("johndoe@gmail.com")

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("failed to store forgot password repository", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "kesalahan saat menyimpan data atur ulang password",
			Status:  500,
		}

		userRepo.Mock.On("FindOneByEmail", mock.Anything).Return(&user.User{Id: mock.Anything}).Once()

		forgotPasswordRepo.Mock.On("Insert", mock.Anything).Return(errors.New(mock.Anything)).Once()

		id.Mock.On("MakeIdentifier").Return(mock.Anything).Once()

		bcrypt.Mock.On("Hash", mock.Anything).Return(mock.Anything).Once()

		res := service.ForgotPassword("johndoe@gmail.com")

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("failed to send email", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "kesalahan saat mengirim email atur ulang password",
			Status:  500,
		}

		userRepo.Mock.On("FindOneByEmail", mock.Anything).Return(&user.User{Id: mock.Anything}).Once()

		forgotPasswordRepo.Mock.On("Insert", mock.Anything).Return(nil).Once()

		id.Mock.On("MakeIdentifier").Return(mock.Anything).Once()

		bcrypt.Mock.On("Hash", mock.Anything).Return(mock.Anything).Once()

		smtpMailer.Mock.On("Send", mock.Anything, mock.Anything, mock.Anything).Return(errors.New(mock.Anything)).Once()

		res := service.ForgotPassword("johndoe@gmail.com")

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("success operation", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "link atur password telah dikirim",
			Success: true,
			Status:  200,
		}

		userRepo.Mock.On("FindOneByEmail", mock.Anything).Return(&user.User{Id: mock.Anything}).Once()

		forgotPasswordRepo.Mock.On("Insert", mock.Anything).Return(nil).Once()

		id.Mock.On("MakeIdentifier").Return(mock.Anything).Once()

		bcrypt.Mock.On("Hash", mock.Anything).Return(mock.Anything).Once()

		smtpMailer.Mock.On("Send", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

		res := service.ForgotPassword("johndoe@gmail.com")

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})
}

func TestGetProfile(t *testing.T) {
	t.Run("empty user id", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "id pengguna tidak boleh kosong",
			Status:  400,
		}

		res := service.GetProfile("")

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("user retrieved", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "profile berhasil diambil",
			Success: true,
			Status:  200,
		}

		userRepo.Mock.On("FindOneById", mock.Anything).Return(&user.User{
			Id:          mock.Anything,
			Name:        mock.Anything,
			CompanyName: mock.Anything,
			Email:       mock.Anything,
			Password:    mock.Anything,
		}).Once()

		res := service.GetProfile(mock.Anything)

		common_testing.Assertion(t, expected, res, &common_testing.Options{DataNotEmpty: true})
	})
}

func TestLogin(t *testing.T) {
	t.Run("invalid input", func(t *testing.T) {
		t.Run("email testing", func(t *testing.T) {
			for _, tb := range emailTestTable {
				res := service.Login(&user.LoginData{Email: tb.email})

				common_testing.Assertion(t, tb.expected, res, common_testing.DefaultOption)
			}
		})

		t.Run("empty password", func(t *testing.T) {
			expected := common_testing.Expectation{
				Message: "password tidak boleh kosong",
				Status:  400,
			}

			res := service.Login(&user.LoginData{Email: "johndoe@gmail.com", Password: ""})

			common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
		})
	})

	t.Run("user not registered", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "email atau password salah",
			Status:  401,
		}

		userRepo.Mock.On("FindOneByEmail", mock.Anything).Return(&user.User{}).Once()

		res := service.Login(&user.LoginData{Email: "johndoe@gmail.com", Password: mock.Anything})

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("wrong password", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "email atau password salah",
			Status:  401,
		}

		userRepo.Mock.On("FindOneByEmail", mock.Anything).Return(&user.User{Id: mock.Anything, Password: mock.Anything}).Once()

		bcrypt.Mock.On("Compare", mock.Anything, mock.Anything).Return(false).Once()

		res := service.Login(&user.LoginData{Email: "johndoe@gmail.com", Password: mock.Anything})

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("failed to create token", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "kesalahan saat membuat token",
			Status:  500,
		}

		userRepo.Mock.On("FindOneByEmail", mock.Anything).Return(&user.User{Id: mock.Anything, Password: mock.Anything}).Once()

		bcrypt.Mock.On("Compare", mock.Anything, mock.Anything).Return(true).Once()

		jsonWebToken.Mock.On("Sign", mock.Anything).Return(mock.Anything, errors.New(mock.Anything)).Once()

		res := service.Login(&user.LoginData{Email: "johndoe@gmail.com", Password: mock.Anything})

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("success operation", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "login berhasil",
			Success: true,
			Status:  200,
		}

		userRepo.Mock.On("FindOneByEmail", mock.Anything).Return(&user.User{Id: mock.Anything, Password: mock.Anything}).Once()

		bcrypt.Mock.On("Compare", mock.Anything, mock.Anything).Return(true).Once()

		jsonWebToken.Mock.On("Sign", mock.Anything).Return(mock.Anything, nil).Once()

		res := service.Login(&user.LoginData{Email: "johndoe@gmail.com", Password: mock.Anything})

		common_testing.Assertion(t, expected, res, &common_testing.Options{DataNotEmpty: true, TokenNotEmpty: true})
	})
}

func TestRegister(t *testing.T) {
	t.Run("invalid input", func(t *testing.T) {
		t.Run("email testing", func(t *testing.T) {
			for _, tb := range emailTestTable {
				t.Run(tb.label, func(t *testing.T) {
					res := service.Register(&user.RegisterData{Email: tb.email})

					common_testing.Assertion(t, tb.expected, res, common_testing.DefaultOption)
				})

			}
		})

		t.Run("other input", func(t *testing.T) {
			table := []struct {
				label    string
				input    user.RegisterData
				expected common_testing.Expectation
			}{
				{
					label: "empty name",
					input: user.RegisterData{
						Name:           "",
						CompanyName:    mock.Anything,
						Email:          "johndoe@gmail.com",
						Password:       mock.Anything,
						CofirmPassword: mock.Anything,
					},
					expected: common_testing.Expectation{
						Message: "nama tidak boleh kosong",
						Status:  400,
					},
				},
				{
					label: "empty company name",
					input: user.RegisterData{
						Name:           mock.Anything,
						CompanyName:    "",
						Email:          "johndoe@gmail.com",
						Password:       mock.Anything,
						CofirmPassword: mock.Anything,
					},
					expected: common_testing.Expectation{
						Message: "nama perusahaan tidak boleh kosong",
						Status:  400,
					},
				},
				{
					label: "empty password",
					input: user.RegisterData{
						Name:           mock.Anything,
						CompanyName:    mock.Anything,
						Email:          "johndoe@gmail.com",
						Password:       "",
						CofirmPassword: mock.Anything,
					},
					expected: common_testing.Expectation{
						Message: "password tidak boleh kosong",
						Status:  400,
					},
				},
				{
					label: "empty confirmation password",
					input: user.RegisterData{
						Name:           mock.Anything,
						CompanyName:    mock.Anything,
						Email:          "johndoe@gmail.com",
						Password:       mock.Anything,
						CofirmPassword: "",
					},
					expected: common_testing.Expectation{
						Message: "konfirmasi password tidak boleh kosong",
						Status:  400,
					},
				},
				{
					label: "length less that required length",
					input: user.RegisterData{
						Name:           mock.Anything,
						CompanyName:    mock.Anything,
						Email:          "johndoe@gmail.com",
						Password:       "abcs",
						CofirmPassword: "abcs",
					},
					expected: common_testing.Expectation{
						Message: "panjang password harus lebih dari 8 karakter",
						Status:  400,
					},
				},
				{
					label: "password not confirmed",
					input: user.RegisterData{
						Name:           mock.Anything,
						CompanyName:    mock.Anything,
						Email:          "johndoe@gmail.com",
						Password:       mock.Anything,
						CofirmPassword: "ubahpassword123",
					},
					expected: common_testing.Expectation{
						Message: "konfirmasi password salah",
						Status:  400,
					},
				},
			}

			for _, tb := range table {
				t.Run(tb.label, func(t *testing.T) {
					res := service.Register(&tb.input)

					common_testing.Assertion(t, tb.expected, res, common_testing.DefaultOption)
				})
			}
		})
	})

	t.Run("not verified", func(t *testing.T) {
		input := user.RegisterData{
			Name:           mock.Anything,
			CompanyName:    mock.Anything,
			Email:          "johndoe@gmail.com",
			Password:       mock.Anything,
			CofirmPassword: mock.Anything,
		}
		expected := common_testing.Expectation{
			Message: "email tidak tervalidasi",
			Status:  401,
		}

		verificationRepo.Mock.On("FindOneByEmail", mock.Anything).Return(&verification.Verification{Status: "not verified"}).Once()

		res := service.Register(&input)

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("user registered", func(t *testing.T) {
		input := user.RegisterData{
			Name:           mock.Anything,
			CompanyName:    mock.Anything,
			Email:          "johndoe@gmail.com",
			Password:       mock.Anything,
			CofirmPassword: mock.Anything,
		}
		expected := common_testing.Expectation{
			Message: "pengguna sudah terdaftar",
			Status:  409,
		}

		verificationRepo.Mock.On("FindOneByEmail", mock.Anything).Return(&verification.Verification{Status: "verified"}).Once()

		userRepo.Mock.On("FindOneByEmail", mock.Anything).Return(&user.User{Id: mock.Anything}).Once()

		res := service.Register(&input)

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("failed to store user", func(t *testing.T) {
		input := user.RegisterData{
			Name:           mock.Anything,
			CompanyName:    mock.Anything,
			Email:          "johndoe@gmail.com",
			Password:       mock.Anything,
			CofirmPassword: mock.Anything,
		}
		expected := common_testing.Expectation{
			Message: "kesalahan saat menyimpan pengguna",
			Status:  500,
		}

		verificationRepo.Mock.On("FindOneByEmail", mock.Anything).Return(&verification.Verification{Status: "verified"}).Once()

		userRepo.Mock.On("FindOneByEmail", mock.Anything).Return(&user.User{}).Once()

		bcrypt.Mock.On("Hash", mock.Anything).Return(mock.Anything).Once()

		userRepo.Mock.On("Insert", mock.Anything).Return(mock.Anything, errors.New(mock.Anything)).Once()

		res := service.Register(&input)

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("failed to generate token", func(t *testing.T) {
		input := user.RegisterData{
			Name:           mock.Anything,
			CompanyName:    mock.Anything,
			Email:          "johndoe@gmail.com",
			Password:       mock.Anything,
			CofirmPassword: mock.Anything,
		}
		expected := common_testing.Expectation{
			Message: "kesalahan saat membuat token",
			Status:  500,
		}

		verificationRepo.Mock.On("FindOneByEmail", mock.Anything).Return(&verification.Verification{Status: "verified"}).Once()

		userRepo.Mock.On("FindOneByEmail", mock.Anything).Return(&user.User{}).Once()

		bcrypt.Mock.On("Hash", mock.Anything).Return(mock.Anything).Once()

		userRepo.Mock.On("Insert", mock.Anything).Return(mock.Anything, nil).Once()

		jsonWebToken.Mock.On("Sign", mock.Anything).Return(mock.Anything, errors.New(mock.Anything)).Once()

		res := service.Register(&input)

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("success operation", func(t *testing.T) {
		input := user.RegisterData{
			Name:           mock.Anything,
			CompanyName:    mock.Anything,
			Email:          "johndoe@gmail.com",
			Password:       mock.Anything,
			CofirmPassword: mock.Anything,
		}
		expected := common_testing.Expectation{
			Message: "registrasi berhasil",
			Success: true,
			Status:  200,
		}

		verificationRepo.Mock.On("FindOneByEmail", mock.Anything).Return(&verification.Verification{Status: "verified"}).Once()

		userRepo.Mock.On("FindOneByEmail", mock.Anything).Return(&user.User{}).Once()

		bcrypt.Mock.On("Hash", mock.Anything).Return(mock.Anything).Once()

		userRepo.Mock.On("Insert", mock.Anything).Return(mock.Anything, nil).Once()

		jsonWebToken.Mock.On("Sign", mock.Anything).Return(mock.Anything, nil).Once()

		res := service.Register(&input)

		common_testing.Assertion(t, expected, res, &common_testing.Options{DataNotEmpty: true, TokenNotEmpty: true})
	})
}

func TestResetPassword(t *testing.T) {
	t.Run("invalid input", func(t *testing.T) {
		table := []struct {
			label    string
			input    user.ResetPasswordData
			expected common_testing.Expectation
		}{
			{
				label: "empty password",
				input: user.ResetPasswordData{
					Token:          mock.Anything,
					Secret:         mock.Anything,
					Password:       "",
					CofirmPassword: mock.Anything,
					UserId:         mock.Anything,
				},
				expected: common_testing.Expectation{
					Message: "password tidak boleh kosong",
					Status:  400,
				},
			},
			{
				label: "empty confirm password",
				input: user.ResetPasswordData{
					Token:          mock.Anything,
					Secret:         mock.Anything,
					Password:       mock.Anything,
					CofirmPassword: "",
					UserId:         mock.Anything,
				},
				expected: common_testing.Expectation{
					Message: "konfirmasi password tidak boleh kosong",
					Status:  400,
				},
			},
			{
				label: "minumin length not meet",
				input: user.ResetPasswordData{
					Token:          mock.Anything,
					Secret:         mock.Anything,
					Password:       "abc",
					CofirmPassword: mock.Anything,
					UserId:         mock.Anything,
				},
				expected: common_testing.Expectation{
					Message: "panjang password harus lebih dari 8 karakter",
					Status:  400,
				},
			},
			{
				label: "wrong confirmation password",
				input: user.ResetPasswordData{
					Token:          mock.Anything,
					Secret:         mock.Anything,
					Password:       mock.Anything,
					CofirmPassword: "admin123123",
					UserId:         mock.Anything,
				},
				expected: common_testing.Expectation{
					Message: "konfirmasi password salah",
					Status:  400,
				},
			},
			{
				label: "empty otp",
				input: user.ResetPasswordData{
					Token:          mock.Anything,
					Secret:         "",
					Password:       mock.Anything,
					CofirmPassword: mock.Anything,
					UserId:         mock.Anything,
				},
				expected: common_testing.Expectation{
					Message: "kode rahasia tidak boleh kosong",
					Status:  400,
				},
			},
			{
				label: "empty reset token",
				input: user.ResetPasswordData{
					Token:          "",
					Secret:         mock.Anything,
					Password:       mock.Anything,
					CofirmPassword: mock.Anything,
					UserId:         mock.Anything,
				},
				expected: common_testing.Expectation{
					Message: "token reset password tidak boleh kosong",
					Status:  400,
				},
			},
		}

		for _, tb := range table {
			t.Run(tb.label, func(t *testing.T) {
				res := service.ResetPassword(&tb.input)

				common_testing.Assertion(t, tb.expected, res, common_testing.DefaultOption)
			})
		}
	})

	t.Run("forgot password data not found", func(t *testing.T) {
		input := user.ResetPasswordData{
			Token:          mock.Anything,
			Secret:         mock.Anything,
			Password:       mock.Anything,
			CofirmPassword: mock.Anything,
			UserId:         mock.Anything,
		}

		expected := common_testing.Expectation{
			Message: "sesi forgot password tidak ada",
			Status:  404,
		}

		forgotPasswordRepo.Mock.On("FindOneByToken", mock.Anything).Return(&forgot_password.ForgotPassword{}).Once()

		res := service.ResetPassword(&input)

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("wrong reset code", func(t *testing.T) {
		input := user.ResetPasswordData{
			Token:          mock.Anything,
			Secret:         mock.Anything,
			Password:       mock.Anything,
			CofirmPassword: mock.Anything,
			UserId:         mock.Anything,
		}

		expected := common_testing.Expectation{
			Message: "kode reset password salah",
			Status:  401,
		}

		forgotPasswordRepo.Mock.On("FindOneByToken", mock.Anything).Return(&forgot_password.ForgotPassword{Id: mock.Anything, Secret: mock.Anything}).Once()

		bcrypt.Mock.On("Compare", mock.Anything, mock.Anything).Return(false).Once()

		res := service.ResetPassword(&input)

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("failed to update password", func(t *testing.T) {
		input := user.ResetPasswordData{
			Token:          mock.Anything,
			Secret:         mock.Anything,
			Password:       mock.Anything,
			CofirmPassword: mock.Anything,
			UserId:         mock.Anything,
		}

		expected := common_testing.Expectation{
			Message: "kesalahan saat mengatur ulang",
			Status:  500,
		}

		forgotPasswordRepo.Mock.On("FindOneByToken", mock.Anything).Return(&forgot_password.ForgotPassword{
			Id:     mock.Anything,
			Token:  mock.Anything,
			UserId: mock.Anything,
			Secret: mock.Anything,
		}).Once()

		bcrypt.Mock.On("Compare", mock.Anything, mock.Anything).Return(true).Once()

		bcrypt.Mock.On("Hash", mock.Anything).Return(mock.Anything).Once()

		userRepo.Mock.On("UpdatePassword", mock.Anything).Return(errors.New(mock.Anything)).Once()

		res := service.ResetPassword(&input)

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("failed to delete forgot password session", func(t *testing.T) {
		input := user.ResetPasswordData{
			Token:          mock.Anything,
			Secret:         mock.Anything,
			Password:       mock.Anything,
			CofirmPassword: mock.Anything,
			UserId:         mock.Anything,
		}

		expected := common_testing.Expectation{
			Message: "kesalhan saat menghapus sesi reset password",
			Status:  500,
		}

		forgotPasswordRepo.Mock.On("FindOneByToken", mock.Anything).Return(&forgot_password.ForgotPassword{
			Id:     mock.Anything,
			Token:  mock.Anything,
			UserId: mock.Anything,
			Secret: mock.Anything,
		}).Once()

		bcrypt.Mock.On("Compare", mock.Anything, mock.Anything).Return(true).Once()

		bcrypt.Mock.On("Hash", mock.Anything).Return(mock.Anything).Once()

		userRepo.Mock.On("UpdatePassword", mock.Anything).Return(nil).Once()

		forgotPasswordRepo.Mock.On("Delete", mock.Anything).Return(errors.New(mock.Anything)).Once()

		res := service.ResetPassword(&input)

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("success operation", func(t *testing.T) {
		input := user.ResetPasswordData{
			Token:          mock.Anything,
			Secret:         mock.Anything,
			Password:       mock.Anything,
			CofirmPassword: mock.Anything,
			UserId:         mock.Anything,
		}

		expected := common_testing.Expectation{
			Message: "password berhasil diatur ulang",
			Success: true,
			Status:  200,
		}

		forgotPasswordRepo.Mock.On("FindOneByToken", mock.Anything).Return(&forgot_password.ForgotPassword{
			Id:     mock.Anything,
			Token:  mock.Anything,
			UserId: mock.Anything,
			Secret: mock.Anything,
		}).Once()

		bcrypt.Mock.On("Compare", mock.Anything, mock.Anything).Return(true).Once()

		bcrypt.Mock.On("Hash", mock.Anything).Return(mock.Anything).Once()

		userRepo.Mock.On("UpdatePassword", mock.Anything).Return(nil).Once()

		forgotPasswordRepo.Mock.On("Delete", mock.Anything).Return(nil).Once()

		res := service.ResetPassword(&input)

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})
}

func TestUpdateProfile(t *testing.T) {
	t.Run("invalid input", func(t *testing.T) {
		table := []struct {
			label    string
			input    user.User
			expected common_testing.Expectation
		}{
			{
				label: "empty name",
				input: user.User{
					Id:          mock.Anything,
					Name:        "",
					CompanyName: mock.Anything,
				},
				expected: common_testing.Expectation{
					Message: "nama tidak boleh kosong",
					Status:  400,
				},
			},
			{
				label: "empty company name",
				input: user.User{
					Id:          mock.Anything,
					Name:        mock.Anything,
					CompanyName: "",
				},
				expected: common_testing.Expectation{
					Message: "nama perusahaan tidak boleh kosong",
					Status:  400,
				},
			},
			{
				label: "empty id",
				input: user.User{
					Id:          "",
					Name:        mock.Anything,
					CompanyName: mock.Anything,
				},
				expected: common_testing.Expectation{
					Message: "id tidak boleh kosong",
					Status:  400,
				},
			},
		}

		for _, tb := range table {
			res := service.UpdateProfile(&tb.input)

			common_testing.Assertion(t, tb.expected, res, common_testing.DefaultOption)
		}
	})

	t.Run("failed to update repository", func(t *testing.T) {
		input := user.User{
			Id:          mock.Anything,
			Name:        mock.Anything,
			CompanyName: mock.Anything,
		}

		expected := common_testing.Expectation{
			Message: "kesalahan saat update profile",
			Status:  500,
		}

		userRepo.Mock.On("Update", mock.Anything).Return(errors.New(mock.Anything)).Once()

		res := service.UpdateProfile(&input)

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("success operation", func(t *testing.T) {
		input := user.User{
			Id:          mock.Anything,
			Name:        mock.Anything,
			CompanyName: mock.Anything,
		}

		expected := common_testing.Expectation{
			Message: "profile berhasil diupdate",
			Success: true,
			Status:  200,
		}

		userRepo.Mock.On("Update", mock.Anything).Return(nil).Once()

		res := service.UpdateProfile(&input)

		common_testing.Assertion(t, expected, res, &common_testing.Options{DataNotEmpty: true})
	})
}

func TestVerificateEmail(t *testing.T) {
	t.Run("invalid input", func(t *testing.T) {
		table := []struct {
			label    string
			input    verification.Verification
			expected common_testing.Expectation
		}{
			{
				label: "empty email",
				input: verification.Verification{
					RequesterEmail: "",
					Secret:         mock.Anything,
				},
				expected: common_testing.Expectation{
					Message: "email pengaju tidak boleh kosong",
					Status:  400,
				},
			},
			{
				label: "empty secret",
				input: verification.Verification{
					RequesterEmail: mock.Anything,
					Secret:         "",
				},
				expected: common_testing.Expectation{
					Message: "kode OTP tidak boleh kosong",
					Status:  400,
				},
			},
		}

		for _, tb := range table {
			res := service.VerificateEmail(&tb.input)

			common_testing.Assertion(t, tb.expected, res, common_testing.DefaultOption)
		}
	})

	t.Run("no verification session found", func(t *testing.T) {
		input := verification.Verification{
			RequesterEmail: mock.Anything,
			Secret:         mock.Anything,
		}
		expected := common_testing.Expectation{
			Message: "sesi verifikasi tidak ada",
			Status:  404,
		}

		verificationRepo.Mock.On("FindOneByEmail", mock.Anything).Return(&verification.Verification{}).Once()

		res := service.VerificateEmail(&input)

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("already verified email", func(t *testing.T) {
		input := verification.Verification{
			RequesterEmail: mock.Anything,
			Secret:         mock.Anything,
		}
		expected := common_testing.Expectation{
			Message: "email sudah terverifikasi",
			Status:  409,
		}

		verificationRepo.Mock.On("FindOneByEmail", mock.Anything).Return(&verification.Verification{Id: mock.Anything, Status: "verified"}).Once()

		res := service.VerificateEmail(&input)

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("wrong otp", func(t *testing.T) {
		input := verification.Verification{
			RequesterEmail: mock.Anything,
			Secret:         mock.Anything,
		}

		expected := common_testing.Expectation{
			Message: "kode verifikasi salah",
			Status:  200,
		}

		verificationRepo.Mock.On("FindOneByEmail", mock.Anything).Return(&verification.Verification{Id: mock.Anything, Status: "not verified"}).Once()

		bcrypt.Mock.On("Compare", mock.Anything, mock.Anything).Return(false).Once()

		res := service.VerificateEmail(&input)

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("already verified email", func(t *testing.T) {
		input := verification.Verification{
			RequesterEmail: mock.Anything,
			Secret:         mock.Anything,
		}

		expected := common_testing.Expectation{
			Message: "kesalahan saat update status verifikasi",
			Status:  500,
		}

		verificationRepo.Mock.On("FindOneByEmail", mock.Anything).Return(&verification.Verification{Id: mock.Anything, Status: "not verified"}).Once()

		bcrypt.Mock.On("Compare", mock.Anything, mock.Anything).Return(true).Once()

		verificationRepo.Mock.On("UpdateStatus", mock.Anything).Return(errors.New(mock.Anything)).Once()

		res := service.VerificateEmail(&input)

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("already verified email", func(t *testing.T) {
		input := verification.Verification{
			RequesterEmail: mock.Anything,
			Secret:         mock.Anything,
		}

		expected := common_testing.Expectation{
			Message: "email terverifikasi",
			Success: true,
			Status:  200,
		}

		verificationRepo.Mock.On("FindOneByEmail", mock.Anything).Return(&verification.Verification{Id: mock.Anything, Status: "not verified"}).Once()

		bcrypt.Mock.On("Compare", mock.Anything, mock.Anything).Return(true).Once()

		verificationRepo.Mock.On("UpdateStatus", mock.Anything).Return(nil).Once()

		res := service.VerificateEmail(&input)

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})
}

func TestVerificationRequest(t *testing.T) {
	t.Run("invalid email", func(t *testing.T) {
		for _, tb := range emailTestTable {
			t.Run(tb.label, func(t *testing.T) {
				res := service.VerificationRequest(tb.email)

				common_testing.Assertion(t, tb.expected, res, common_testing.DefaultOption)
			})
		}
	})

	t.Run("already verification", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "email sudah terverifikasi",
			Status:  409,
		}

		verificationRepo.Mock.On("FindOneByEmail", mock.Anything).Return(&verification.Verification{Id: mock.Anything, Status: "verified"}).Once()

		res := service.VerificationRequest("johndoe@gmail.com")

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("verification not found", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "kesalahan saat menyimpan sesi verifikasi",
			Status:  500,
		}

		verificationRepo.Mock.On("FindOneByEmail", mock.Anything).Return(&verification.Verification{Id: mock.Anything, Status: "not verified"}).Once()

		bcrypt.Mock.On("Hash", mock.Anything).Return(mock.Anything).Once()

		verificationRepo.Mock.On("Upsert", mock.Anything).Return(errors.New(mock.Anything)).Once()

		res := service.VerificationRequest("johndoe@gmail.com")

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("failed to send email", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "kesalahan saat mengirim email",
			Status:  500,
		}

		verificationRepo.Mock.On("FindOneByEmail", mock.Anything).Return(&verification.Verification{Id: mock.Anything, Status: "not verified"}).Once()

		bcrypt.Mock.On("Hash", mock.Anything).Return(mock.Anything).Once()

		verificationRepo.Mock.On("Upsert", mock.Anything).Return(nil).Once()

		smtpMailer.Mock.On("Send", mock.Anything, mock.Anything, mock.Anything).Return(errors.New(mock.Anything)).Once()

		res := service.VerificationRequest("johndoe@gmail.com")

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("success operation", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "verifikasi email berhasil",
			Success: true,
			Status:  200,
		}

		verificationRepo.Mock.On("FindOneByEmail", mock.Anything).Return(&verification.Verification{Id: mock.Anything, Status: "not verified"}).Once()

		bcrypt.Mock.On("Hash", mock.Anything).Return(mock.Anything).Once()

		verificationRepo.Mock.On("Upsert", mock.Anything).Return(nil).Once()

		smtpMailer.Mock.On("Send", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

		res := service.VerificationRequest("johndoe@gmail.com")

		common_testing.Assertion(t, expected, res, &common_testing.Options{DataNotEmpty: true})
	})
}
