package domain

type Credentials struct {
	Email    string
	Password string
}

type JWT struct {
	AccessToken string
}
