package fake

type AuthFake struct{}

func (*AuthFake) Authenticate(username string, password string) bool {
	return username == "user" && password == "good password"
}
