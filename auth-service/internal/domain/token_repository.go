package domain

type TokenRepository interface {
	GenerateAccessToken(userID string) (string, error)
	VerifyAccessToken(tokenString string) (string, error) // Devuelve userID si es válido
	GenerateRefreshToken() (string, error)
}