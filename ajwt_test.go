package ajwt

import (
	"testing"
	"time"
)

func TestGenerateJWT(t *testing.T) {
	claims := make(map[string]interface{})
	claims["name"] = "username"
	s, err := GenerateJWT(time.Hour*3, []byte("signingKey"), claims)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
	if len(s) == 0 {
		t.Log("token should have non-zero length", err)
		t.Fail()
	}
}
