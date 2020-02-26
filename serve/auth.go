package serve

import (
	"app/database"
	"app/util/random"
	"errors"
	"github.com/gin-gonic/gin"
)

// generate a new refresh token
func NewRefrshToken(c gin.Context, userName string, uid uint64) (string, error) {

	rereshToken, err := random.GetRandomString(64)
	if err != nil {
		return "", err
	}

	if err := database.CreateToken(database.RefreshTokenTable, rereshToken, uid); err != nil {
		return "", err
	}

	return rereshToken, nil
}

// generate a new access token
func NewAccessToken(refreshToken string) (string, error) {
	// check refresh token first
	uid, err := GetuidByToken(database.RefreshTokenTable, refreshToken)
	if err != nil {
		return "", err
	}

	accessToken, err := hashToken(refreshToken)
	if err != nil {
		return "", err
	}

	if err := database.CreateToken(database.AccessTokenTable, accessToken, uid); err != nil {
		return "", err
	}

	return accessToken, nil
}

// refresh token uid and uid not same error
var ERR_INVALID_USER error = errors.New("Invalid of user with this refresh token")

// check a refresh token is valid
func GetuidByRefreshToken(refreshToken string) (uint64, error) {
	return GetuidByToken(database.RefreshTokenTable, refreshToken)
}

// check a access token is valid
func GetuidByAccessToken(accessToken string) (uint64, error) {
	return GetuidByToken(database.AccessTokenTable, accessToken)
}

// check a token is valid
func GetuidByToken(authType, token string) (uint64, error) {
	result, err := database.GetuidbyToken(authType, token)
	if err != nil {
		return 0, err
	}

	return result, nil
}
