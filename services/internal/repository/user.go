package repository

import (
	"database/sql"
	"fmt"
	"go-web-robotek/services/internal/domain"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

var secretKey = []byte("secretkey")

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (r *UserRepo) SignIn(c *domain.Credentials) (*domain.JWT, error) {
	stmt := `SELECT Password FROM users WHERE email = $1`

	var password string
	err := r.db.QueryRow(stmt, c.Email).Scan(&password)
	if err != nil {
		return nil, err
	}

	if password != c.Password {
		return nil, fmt.Errorf("wrong password")
	}

	tokenString, err := createToken(c.Email)
	if err != nil {
		return nil, err
	}

	return &domain.JWT{AccessToken: tokenString}, nil
}
