package auth

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	errors "github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	models "github.com/krezac/los-server/models"
)

type JwtExtraOptions struct {
	JwtSigningKey string `long:"jwt-signing-key" description:"Key to sign JWT tokens" env:"JWT_SIGNING_KEY" required:"true"`
	JwtVerifyKey  string `long:"jwt-verify-key" description:"Key to verify JWT tokens" env:"JWT_VERIFY_KEY" required:"true"`
	JwtIssuerName string `long:"jwt-issuer-name" description:"Issuer name for JWT tokens" env:"JWT_ISSUER_NAME" required:"true"`
}

var JwtExtraOptionsVar = &JwtExtraOptions{}

// code from github.com/go-swagger/go-swagger/examples/composed-auth/auth/authorizers.go

var (
	// Keys used to sign and verify our tokens
	verifyKey *rsa.PublicKey
	// currently unused: signKey   *rsa.PrivateKey
)

// roleClaims describes the format of our JWT token's claims
type roleClaims struct {
	Roles []string `json:"roles"`
	jwt.StandardClaims
}

// Customized authorizer methods for our sample API

// HasRole tells if the Bearer token is a JWT signed by us with a claim to be
// member of an authorization scope.
// We verify that the claimed role is one of the passed scopes
func HasRole(token string, scopes []string) (*models.Principal, error) {
	claims, err := parseAndCheckToken(token)
	if err == nil {
		if claims.Issuer == JwtExtraOptionsVar.JwtIssuerName {
			isInScopes := false
			claimedRoles := []string{}
			for _, scope := range scopes {
				for _, role := range claims.Roles {
					if role == scope {
						isInScopes = true
						// we enrich the principal with all claimed roles within scope (hence: not breaking here)
						claimedRoles = append(claimedRoles, role)
					}
				}
			}
			if isInScopes {
				return &models.Principal{
					Name:     claims.Id,
					Roles:    claimedRoles,
					RawToken: token,
					ValidTo:  strfmt.DateTime(time.Unix(claims.ExpiresAt, 0)),
				}, nil
			}
			return nil, errors.New(403, "Forbidden: insufficient privileges")
		}
	}
	return nil, errors.New(401, "Unauthorized: invalid Bearer token: %v", err)
}

func parseAndCheckToken(token string) (*roleClaims, error) {
	if verifyKey == nil {
		// loads public keys to verify our tokens
		verifyKeyBuf, err := ioutil.ReadFile(JwtExtraOptionsVar.JwtVerifyKey)
		if err != nil {
			fmt.Printf("%#v", JwtExtraOptionsVar)
			panic("Cannot load public key for tokens " + JwtExtraOptionsVar.JwtVerifyKey + " " + err.Error())
		}
		verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyKeyBuf)
		if err != nil {
			panic("Invalid public key for tokens " + err.Error())
		}

	}
	// the API key is a JWT signed by us with a claim to be a reseller
	parsedToken, err := jwt.ParseWithClaims(token, &roleClaims{}, func(parsedToken *jwt.Token) (interface{}, error) {
		// the key used to validate tokens
		return verifyKey, nil
	})

	if err == nil {
		if claims, ok := parsedToken.Claims.(*roleClaims); ok && parsedToken.Valid {
			return claims, nil
		}
	}
	return nil, err

}
