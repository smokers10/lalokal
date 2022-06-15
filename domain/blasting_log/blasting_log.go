package blasting_log

var Status = []string{"sent", "failed"}

type BlastingLogDomain struct {
	Id                string `json:"id" form:"id" bson:"_id"`
	Status            string `json:"status" form:"status" bson:"status"`
	BlastingSessionId string `json:"blasting_session_id" form:"blasting_session_id" bson:"blasting_session_id"`
	TopicId           string `json:"topic_id" form:"topic_id" bson:"topic_id"`
}

type Repository interface {
	Insert(data *BlastingLogDomain) (failure error)

	FindByTopicId(topic_id string) (result []BlastingLogDomain)

	Count(topic_id string) (count int)
}
