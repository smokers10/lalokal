package encryption

type Contract interface {
	Hash(plaintext string) (hashed_string string)

	Compare(hashed_text string, plain_text string) (is_correct bool)
}
