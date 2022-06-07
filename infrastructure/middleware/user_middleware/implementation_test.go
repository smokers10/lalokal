package user_middleware

import (
	"errors"
	"lalokal/domain/user"
	"lalokal/domain/verification"
	"lalokal/infrastructure/jsonwebtoken"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userRepository         = user.MockRepository{Mock: mock.Mock{}}
	verificationRepository = verification.MockRepository{Mock: mock.Mock{}}
	jwt                    = jsonwebtoken.MockContract{Mock: mock.Mock{}}
	contractImpl           = userMiddleware{
		userRepository:         &userRepository,
		verificationRepository: &verificationRepository,
		jwt:                    &jwt,
	}
	commonTesting = func(t *testing.T, expected middlewareResult, actual middlewareResult, enable_claims_assert bool) {
		assert.Equal(t, expected.Message, actual.Message)
		assert.Equal(t, expected.Status, actual.Status)
		assert.Equal(t, expected.Reason, actual.Reason)

		if enable_claims_assert {
			assert.NotEmpty(t, actual.Claim)
		}
	}
)

func TestProcess(t *testing.T) {
	t.Run("empty token", func(t *testing.T) {
		expected := middlewareResult{
			Message: "silahkan login terlebih dahulu",
			Status:  401,
			Reason:  "empty token",
		}

		res := contractImpl.Process("")

		commonTesting(t, expected, *res, false)
	})

	t.Run("empty token", func(t *testing.T) {
		expected := middlewareResult{
			Message: "kesalahan saat parsing token",
			Status:  500,
			Reason:  "error",
		}

		jwt.Mock.On("ParseToken", mock.Anything).Return(map[string]interface{}{}, errors.New(mock.Anything)).Once()

		res := contractImpl.Process(mock.Anything)

		commonTesting(t, expected, *res, false)
	})

	t.Run("user not found", func(t *testing.T) {
		expected := middlewareResult{
			Message: "pengguna tidak terdaftar",
			Status:  401,
			Reason:  "unregistered",
		}

		jwt.Mock.On("ParseToken", mock.Anything).Return(map[string]interface{}{
			"id":    mock.Anything,
			"email": mock.Anything,
		}, nil).Once()

		userRepository.Mock.On("FindOneById", mock.Anything).Return(&user.User{}).Once()

		res := contractImpl.Process(mock.Anything)

		commonTesting(t, expected, *res, false)
	})

	t.Run("verification (not found)", func(t *testing.T) {
		expected := middlewareResult{
			Message: "pengguna tidak terverifikasi",
			Status:  401,
			Reason:  "not verified",
		}

		jwt.Mock.On("ParseToken", mock.Anything).Return(map[string]interface{}{
			"id":    mock.Anything,
			"email": mock.Anything,
		}, nil).Once()

		userRepository.Mock.On("FindOneById", mock.Anything).Return(&user.User{Email: mock.Anything}).Once()

		verificationRepository.Mock.On("FindOneByEmail", mock.Anything).Return(&verification.Verification{}).Once()

		res := contractImpl.Process(mock.Anything)

		commonTesting(t, expected, *res, false)
	})

	t.Run("verification (not verified)", func(t *testing.T) {
		expected := middlewareResult{
			Message: "pengguna tidak terverifikasi",
			Status:  401,
			Reason:  "not verified",
		}

		jwt.Mock.On("ParseToken", mock.Anything).Return(map[string]interface{}{
			"id":    mock.Anything,
			"email": mock.Anything,
		}, nil).Once()

		userRepository.Mock.On("FindOneById", mock.Anything).Return(&user.User{Email: mock.Anything}).Once()

		verificationRepository.Mock.On("FindOneByEmail", mock.Anything).Return(&verification.Verification{Id: mock.Anything, Status: "not verified"}).Once()

		res := contractImpl.Process(mock.Anything)

		commonTesting(t, expected, *res, false)
	})

	t.Run("all pass", func(t *testing.T) {
		expected := middlewareResult{
			Status:  200,
			Is_pass: true,
			Reason:  "all pass",
			Message: "ok",
			Claim: struct {
				Id    string "json:\"id,omitempty\""
				Email string "json:\"email,omitempty\""
			}{
				Email: mock.Anything,
				Id:    mock.Anything,
			},
		}

		jwt.Mock.On("ParseToken", mock.Anything).Return(map[string]interface{}{
			"id":    mock.Anything,
			"email": mock.Anything,
		}, nil).Once()

		userRepository.Mock.On("FindOneById", mock.Anything).Return(&user.User{Email: mock.Anything}).Once()

		verificationRepository.Mock.On("FindOneByEmail", mock.Anything).Return(&verification.Verification{Id: mock.Anything, Status: "verified"}).Once()

		res := contractImpl.Process(mock.Anything)

		commonTesting(t, expected, *res, false)
	})
}
