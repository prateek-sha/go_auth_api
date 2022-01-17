package access_token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpired(), "Brand new access token should not be expired.")
	assert.True(t, at.UserId == 0, "new user should not have any id")
}
