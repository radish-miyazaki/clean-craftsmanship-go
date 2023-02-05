package dummy

type Auth interface {
	Authenticate(username string, password string) bool
}

type AuthDummy struct{}

func (*AuthDummy) Authenticate(_ string, _ string) bool {
	return false
}
