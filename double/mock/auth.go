package mock

type AuthMock struct {
	count            int
	result           bool
	lastUsername     string
	lastPassword     string
	expectedUsername string
	expectedPassword string
	expectedCount    int
}

func NewAuthMock(username string, password string, count int) *AuthMock {
	return &AuthMock{
		expectedUsername: username,
		expectedPassword: password,
		expectedCount:    count,
	}
}

func (auth *AuthMock) Validate() bool {
	return auth.count == auth.expectedCount &&
		auth.lastUsername == auth.expectedUsername &&
		auth.lastPassword == auth.expectedPassword
}

func (auth *AuthMock) Authenticate(username string, password string) bool {
	auth.count++
	auth.lastPassword = password
	auth.lastUsername = username
	return auth.result
}

func (auth *AuthMock) SetResult(result bool) {
	auth.result = result
}

func (auth *AuthMock) GetCount() int {
	return auth.count
}

func (auth *AuthMock) GetLastUsername() string {
	return auth.lastUsername
}

func (auth *AuthMock) GetLastPassword() string {
	return auth.lastPassword
}
