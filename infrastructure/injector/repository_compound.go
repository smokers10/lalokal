package injector

import (
	"lalokal/domain/blasting_log"
	"lalokal/domain/blasting_session"
	"lalokal/domain/forgot_password"
	"lalokal/domain/keyword"
	"lalokal/domain/topic"
	"lalokal/domain/twitter_api_token"
	"lalokal/domain/user"
	"lalokal/domain/verification"
	"lalokal/infrastructure/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

type repositoryCompund struct {
	BlastLogRepository        blasting_log.Repository
	BlastingSessionRepository blasting_session.Repository
	ForgotPasswordRepository  forgot_password.Repository
	KeywordRepository         keyword.Repository
	TopicRepository           topic.Repository
	TwitterAPITokenRepository twitter_api_token.Repository
	UserRepository            user.Repository
	VerificationRepository    verification.Repository
}

func repoCompound(db *mongo.Database) repositoryCompund {
	return repositoryCompund{
		BlastLogRepository:        repository.BlastingLogRepository(db),
		BlastingSessionRepository: repository.BlastingSessionRepository(db),
		ForgotPasswordRepository:  repository.ForgotPasswordRepository(db),
		KeywordRepository:         repository.KeywordRepository(db),
		TopicRepository:           repository.TopicRepository(db),
		TwitterAPITokenRepository: repository.TwitterAPITokenRepository(db),
		UserRepository:            repository.UserRepository(db),
		VerificationRepository:    repository.VerificationRepository(db),
	}
}
