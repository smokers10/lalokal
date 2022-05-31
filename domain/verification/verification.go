package verification

type Verification struct {
	Id             string `json:"id" form:"id" bson:"_id"`
	RequesterEmail string `json:"requester_email" form:"requester_email" bson:"requester_email"`
	Secret         string `json:"secret" form:"secret" bson:"secret"`
	Status         string `json:"status" form:"status" bson:"status"`
}

type Repository interface {
	Upsert(data *Verification) (failure error)

	UpdateStatus(verification_id string) (failure error)

	FindOneByEmail(email string) (result *Verification)
}
