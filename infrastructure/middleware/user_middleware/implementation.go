package user_middleware

import (
	"lalokal/domain/user"
	"lalokal/domain/verification"
	"lalokal/infrastructure/jsonwebtoken"
)

type userMiddleware struct {
	userRepository         user.Repository
	verificationRepository verification.Repository
	jwt                    jsonwebtoken.Contact
}

func UserMiddleware(ur *user.Repository, vr *verification.Repository, jwt *jsonwebtoken.Contact) contract {
	return &userMiddleware{
		userRepository:         *ur,
		verificationRepository: *vr,
		jwt:                    *jwt,
	}
}

func (um *userMiddleware) Process(token string) *middlewareResult {
	// if token are empty
	if token == "" {
		return &middlewareResult{
			Message: "silahkan login terlebih dahulu",
			Status:  401,
			Reason:  "empty token",
		}
	}

	// parsing payload
	payload, err := um.jwt.ParseToken(token)
	if err != nil {
		return &middlewareResult{
			Message: "kesalahan saat parsing token",
			Status:  500,
			Reason:  "error",
		}
	}

	id := payload["id"].(string)
	email := payload["email"].(string)

	// retrieve user by its id
	user := um.userRepository.FindOneById(id)

	// if user not exists
	if user.Email == "" {
		return &middlewareResult{
			Message: "pengguna tidak terdaftar",
			Status:  401,
			Reason:  "unregistered",
		}
	}

	// retrieve verification
	verif := um.verificationRepository.FindOneByEmail(email)

	// if verification not found or not verified
	if verif.Id == "" || verif.Status == "not verified" {
		return &middlewareResult{
			Message: "pengguna tidak terverifikasi",
			Status:  401,
			Reason:  "not verified",
		}
	}

	return &middlewareResult{
		Status:  200,
		Is_pass: true,
		Reason:  "all pass",
		Message: "ok",
		Claim: struct {
			Id    string "json:\"id,omitempty\""
			Email string "json:\"email,omitempty\""
		}{
			Id:    id,
			Email: email,
		},
	}
}
