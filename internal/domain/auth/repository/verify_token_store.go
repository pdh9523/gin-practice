package repository

type VerifyTokenStore interface {
	Save(token, email string) error
	FindEmailByToken(token string) (string, error)
	Delete(token string) error
}
