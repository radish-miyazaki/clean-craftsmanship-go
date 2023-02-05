package spy

type AuthSpy struct {
	count        int
	result       bool
	lastUsername string
	lastPassword string
}

func (auth *AuthSpy) Authenticate(username string, password string) bool {
	auth.count++
	auth.lastPassword = password
	auth.lastUsername = username
	return auth.result
}

func (auth *AuthSpy) SetResult(result bool) {
	auth.result = result
}

func (auth *AuthSpy) GetCount() int {
	return auth.count
}

func (auth *AuthSpy) GetLastUsername() string {
	return auth.lastUsername
}

func (auth *AuthSpy) GetLastPassword() string {
	return auth.lastPassword
}
