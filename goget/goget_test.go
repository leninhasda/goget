package goget

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetMd5Hash(t *testing.T) {
	getter := New(http.DefaultClient)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, client"))
	}))
	defer ts.Close()

	md5, err := getter.GetMd5Hash(ts.URL)
	if err != nil {
		t.Errorf("error: %s", err.Error())
	}
	if md5 != "b819c21b48cb2fa83e9ec2be9630ae9f" {
		t.Errorf("test failed: expected b819c21b48cb2fa83e9ec2be9630ae9f, got %s", md5)
	}
}
