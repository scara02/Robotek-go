package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"go-web-robotek/services/internal/domain"
	"go-web-robotek/services/internal/usecase"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("secretkey")

type UserDelivery struct {
	userUseCase usecase.User
}

func NewUserDelivery(userUseCase usecase.User) *UserDelivery {
	return &UserDelivery{
		userUseCase: userUseCase,
	}
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func (d *UserDelivery) SignIn(w http.ResponseWriter, r *http.Request) {
	var c domain.Credentials

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	accessToken, err := d.userUseCase.SignIn(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	accessTokenJSON, err := json.Marshal(accessToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(string(accessTokenJSON)))
}

func (d *UserDelivery) IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.Split(r.Header.Get("Authorization"), "Bearer ")

		if len(token) == 1 {
			http.Error(w, "no access token provided", http.StatusForbidden)
			return
		}

		err := verifyToken(token[1])
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
