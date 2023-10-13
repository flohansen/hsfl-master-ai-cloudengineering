package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJwtAuthorizer(t *testing.T) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 4096)
	publicKey := &privateKey.PublicKey
	tokenGenerator := JwtTokenGenerator{privateKey, publicKey}

	t.Run("CreateToken", func(t *testing.T) {
		t.Run("should generate valid JWT token", func(t *testing.T) {
			// given
			// when
			token, err := tokenGenerator.CreateToken(map[string]interface{}{
				"exp":  12345,
				"user": "test",
			})

			// then
			assert.NoError(t, err)
			tokenParts := strings.Split(token, ".")
			assert.Len(t, tokenParts, 3)

			b, _ := base64.
				StdEncoding.
				WithPadding(base64.NoPadding).
				DecodeString(tokenParts[1])

			var claims map[string]interface{}
			json.Unmarshal(b, &claims)

			assert.Equal(t, float64(12345), claims["exp"])
			assert.Equal(t, "test", claims["user"])
		})
	})
}
