package stub

type Auth interface {
	Authenticate(username string, password string) bool
}

type RejectingAuth struct{}

func (*RejectingAuth) Authenticate(_ string, _ string) bool {
	return false
}

type PromiscuousAuth struct{}

func (*PromiscuousAuth) Authenticate(_ string, _ string) bool {
	return true
}
