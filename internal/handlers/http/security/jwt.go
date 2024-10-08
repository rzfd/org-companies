package security

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userID string, jwtSecret string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func ValidateToken(tokenString, secret string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return secret sebagai byte slice
		return []byte(secret), nil
	})

	// Handle error jika ada masalah dalam parsing token
	if err != nil {
		return nil, err
	}

	// Pastikan token valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Return claims
	return claims, nil
}

func JWTMiddleware(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing or malformed JWT")
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")

			claims, err := ValidateToken(token, secret)
			if err != nil {
				fmt.Printf("Error: %+v\n", err)
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired JWT")
			}

			c.Set("user", claims)

			return next(c)
		}
	}
}

func ExtractUserID(c echo.Context) (string, error) {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return "", errors.New("missing authorization header")
	}

	if len(token) < 7 || strings.ToLower(token[:7]) != "bearer " {
		return "", fmt.Errorf("invalid token format")
	}
	token = token[7:]

	jwtSecret := os.Getenv("JWT_TOKEN")
	if jwtSecret == "" {
		return "", errors.New("JWT secret is not set in environment variables at Extract ID")
	}

	claims := &jwt.MapClaims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		fmt.Println("Error parsing token:", err)
		return "", err
	}

	if claims, ok := tkn.Claims.(*jwt.MapClaims); ok && tkn.Valid {
		userID, ok := (*claims)["user_id"].(string)
		if !ok {
			return "", errors.New("user ID not found in token claims")
		}
		return userID, nil
	}

	return "", errors.New("invalid token")
}
