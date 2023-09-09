package utils

import (
	"os"
	"time"
	"zregistry_service/literals"
	"zregistry_service/model"

	"github.com/dgrijalva/jwt-go"
)

// GenerateToken generates a new JWT token with the user's claims.
func GenerateToken(tenant model.TenantMaster) (string, error) {
	type Contacts []model.Contact
	var contacts Contacts
	contacts = tenant.Users[0].Contacts
	// Your JWT secret key should be stored securely.
	secretKey := []byte(os.Getenv("SECRETKEY"))

	// Create a new token with the user's claims
	claims := jwt.MapClaims{
		literals.USERID:   tenant.Users[0].Name,
		literals.EXPIRY:   time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		literals.USERTYPE: tenant.Users[0].UserType,
		literals.CONTACTS: contacts,
		literals.COMPANY:  tenant.Name,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
