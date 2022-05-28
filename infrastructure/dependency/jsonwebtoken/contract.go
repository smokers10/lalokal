package jsonwebtoken

type JsonWebToken interface {
	Sign(payload map[string]interface{}) (token string, failure error)

	ParseToken(token string) (payload map[string]interface{}, failure error)
}
