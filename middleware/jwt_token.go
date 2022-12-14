package middleware

import (
	"attendance_api/models"

	"github.com/golang-jwt/jwt/v4"
)

const studentPrivateKey = "student-private-key"
const adminPrivateKey = "admin-private-key"

func GenrateStudentToken(claims *models.Student) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(studentPrivateKey))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyStudentToken(tokenString string) (bool, *models.Student) {
	claims := &models.Student{}

	token, _ := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(studentPrivateKey), nil
	})

	if token.Valid {
		return true, claims
	}
	return false, claims
}

func GenrateAdminToken(claims *models.Admin) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(adminPrivateKey))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyAdminToken(tokenString string) (bool, *models.Admin) {
	claims := &models.Admin{}

	token, _ := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(adminPrivateKey), nil
	})

	if token.Valid {
		return true, claims
	}
	return false, claims
}
