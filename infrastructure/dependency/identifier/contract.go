package identifier

type Contract interface {
	MakeIdentifier() (id string)
}
