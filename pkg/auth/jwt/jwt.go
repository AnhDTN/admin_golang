package jwt

const defaultKey = "admin_golang_token"
const defaultRefreshKey = "admin_golang_refresh_token"

type JWTService interface {
	GenerateToken(userId string) (string, error)
}

type jwtService struct {
}
