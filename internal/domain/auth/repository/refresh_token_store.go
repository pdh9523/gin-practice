package repository

type RefreshTokenStore interface {
	FindByID(userID uint) (string, error)
	Save(userID uint, value string) error
	Delete(userID uint) error
}
