package jwt

type AuthClaims struct {
	UserID uint
	Email  string
	Role   string
}
