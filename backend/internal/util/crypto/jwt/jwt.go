package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/marcustut/fyp/backend/config"
)

const issuer string = "fyp-auth-api"
const expiryDuration time.Duration = time.Minute * 15

// NewJWTClaimsInput is the input of NewJWTClaims.
type NewJWTClaimsInput struct {
	// userID of the JWT's owner.
	ID string

	// username of the JWT's owner.
	Username string

	// email of the JWT's owner.
	Email string

	// duration until the JWT expires.
	ExpiryDuration *time.Duration
}

// NewJWTClaimsOutput is the output of NewJWTClaims.
type NewJWTClaimsOutput struct {
	// a string containing the encoded JWT claims.
	Token string

	// expiry date of the token.
	ExpiredAt time.Time
}

// NewJWTClaims generate JWT claims for a particular user.
// Defaults to 15min for expiry duration if not specified.
func NewJWTClaims(input *NewJWTClaimsInput) (*NewJWTClaimsOutput, error) {
	// jwt expired time
	var exp time.Time
	// calculate expiry date
	if input.ExpiryDuration != nil {
		exp = time.Now().Add(*input.ExpiryDuration)
	} else {
		exp = time.Now().Add(expiryDuration)
	}

	// generate jwt claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"aud": map[string]interface{}{
			"id":       input.ID,
			"username": input.Username,
			"email":    input.Email,
		},
		"iss": issuer,
		"iat": time.Now().Unix(),
		"exp": exp.Unix(),
	})
	// sign jwt claims with secret key
	tokenString, err := token.SignedString([]byte(config.C.Services.Auth.SecretKey))
	if err != nil {
		return nil, err
	}

	return &NewJWTClaimsOutput{Token: tokenString, ExpiredAt: exp}, nil
}

// ValidateJWTToken takes a signed JWT token string, parses it
// and check for its validity based on expiry date, signing
// method and secret key.
func ValidateJWTToken(token string) (bool, error) {
	// parse token into jwt object
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.C.Services.Auth.SecretKey), nil
	})
	if err != nil {
		return false, err
	}

	// validate the token
	if !t.Valid {
		return false, nil
	}

	return true, nil
}
