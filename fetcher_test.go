package auth0

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"strconv"
	"time"

	"net/http"
	"testing"
)

func TestNewClientCredentialsFlowExpectsObjectInitializedSuccessfully(t *testing.T) {
	// arrange
	clientID := "PV2AvGcMjOFErV6QpaqKnfrUdt8yPuHI"
	clientSecret := "9oXvXvWHfQaaAiWr-wBfS5Vtyp3aGyMuwIwqYs2NuRtmV7-1XEXXNJ1ZA97jLo6J"
	tokenURL := "https://yourcompany.auth0.com/oauth/token"
	grantType := "client_credentials"

	// act
	f := NewTokenFetcher(nil, tokenURL, clientID, clientSecret)

	// assert
	a := f.(*tokenFetcher)
	assert.Equal(t, clientID, a.clientID, "expected client id to match")
	assert.Equal(t, clientSecret, a.clientSecret, "expected client secret to match")
	assert.Equal(t, tokenURL, a.tokenURL, "expected tokenURL to match")
	assert.Equal(t, grantType, a.grantType, "expected grant type to match")
}

func TestConfirmsTokenNotCaching(t *testing.T) {
	i := 1
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{\"access_token\":\"token"+strconv.Itoa(i)+"\"}")
		i++
	}))
	defer ts.Close()

	// arrange
	clientID := "PV2AvGcMjOFErV6QpaqKnfrUdt8yPuHI"
	clientSecret := "9oXvXvWHfQaaAiWr-wBfS5Vtyp3aGyMuwIwqYs2NuRtmV7-1XEXXNJ1ZA97jLo6J"

	// act
	f := NewTokenFetcher(ts.Client(), ts.URL, clientID, clientSecret)
	token, err := f.Token("audience")
	newToken, err := f.NewToken("audience")

	// assert
	assert.Nil(t, err, "Expected no error")
	assert.Equal(t, "token1", token, "Expected 1st token")
	assert.Equal(t, "token2", newToken, "Expected 2nd token")
}

func TestConfirmsTokenCaching(t *testing.T) {
	i := 1
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expiresAt := time.Now().Add(time.Second * 2).Unix()

		token := jwt.New(jwt.SigningMethodHS256)

		token.Claims = &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		}
		tokenString, _ := token.SignedString([]byte("secret"))

		fmt.Fprintln(w, "{\"access_token\":\""+tokenString+"\"}")
		i++
	}))
	defer ts.Close()

	// arrange
	clientID := "PV2AvGcMjOFErV6QpaqKnfrUdt8yPuHI"
	clientSecret := "9oXvXvWHfQaaAiWr-wBfS5Vtyp3aGyMuwIwqYs2NuRtmV7-1XEXXNJ1ZA97jLo6J"

	// act
	f := NewTokenFetcher(ts.Client(), ts.URL, clientID, clientSecret)
	token, err := f.Token("audience")
	token2, err := f.Token("audience")
	time.Sleep(3 * time.Second)
	token3, err := f.Token("audience")

	// assert
	assert.Nil(t, err, "Expected no error")
	assert.Equal(t, token2, token, "Expected tokens to be equal")
	assert.NotEqual(t, token3, token, "Expected new token")
}

// Integration test
func _TestAccessTokenFromGlobalCredentials(t *testing.T) {
	// arrange
	clientID := "zO2mKgrhEA6kcI23E0lRHutHBd1AX8ht"
	clientSecret := "W20FgCL0FZ_ZMFmpdrh13Y49WuflWrJswPqJjDtaXtcaUAYD2x0ETsiPQ1xh8xez"
	tokenURL := "https://yourcompany.auth0.com/oauth/token"

	// act
	f := NewTokenFetcher(http.DefaultClient, tokenURL, clientID, clientSecret)
	accessToken, err := f.NewToken("")

	// assert
	assert.Nil(t, err, "Expected no errors")
	assert.NotEmpty(t, accessToken, "Expected a non-empty access token")
}
