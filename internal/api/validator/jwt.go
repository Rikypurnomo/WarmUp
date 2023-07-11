package validator

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	authdto "github.com/Rikypurnomo/warmup/internal/api/dto/auth"
	"github.com/Rikypurnomo/warmup/internal/api/models"
	"github.com/Rikypurnomo/warmup/pkg/cache"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func GenerateToken(email string, userId int) (authdto.LoginTokenResponse, error) {
	var res authdto.LoginTokenResponse
	res.NewUuid = uuid.New().String()
	res.Expired = time.Now().Add(time.Hour * 24).Unix()
	claims := jwt.MapClaims{
		"id":    userId,
		"uuid":  res.NewUuid,
		"exp":   res.Expired,
		"email": email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("rahasia"))
	if err != nil {
		return res, err
	}

	res.Token = signedToken

	return res, nil

}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte("rahasia"), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func ExtractUserIDFromToken(token *jwt.Token) (uint, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("invalid token claims")
	}

	userID, ok := claims["id"].(float64)
	if !ok {
		return 0, fmt.Errorf("invalid user ID in token claims")
	}

	return uint(userID), nil
}

func CreateAuth(ctx context.Context, customerId uint, token *authdto.LoginTokenResponse, user *models.User) error {

	errAccess := cache.SetKey(ctx, token.NewUuid, customerId, time.Minute*24).Err()

	if errAccess != nil {
		return errAccess
	}
	mars, _ := json.Marshal(user)
	errUser := cache.SetKey(ctx, fmt.Sprintf("user:%s", token.NewUuid), string(mars), time.Minute*5).Err()

	if errUser != nil {
		return errUser
	}

	return nil

}

func ExtractAuth(ctx context.Context, token *authdto.LoginTokenResponse) (string, error) {
	access, err := cache.GetKey(ctx, token.NewUuid).Result()
	if err != nil {
		return "", err
	}

	return access, nil

}
