package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestGetApiKey(t *testing.T) {

	var noAuthHeader = make(http.Header)
	var emptyAuthHeader = make(http.Header)
	var malformedTagAuthHeader = make(http.Header)
	var malformedLengthAuthHeader = make(http.Header)
	//var noKeyAuthHeader = make(http.Header)
	var goodKeyAuthHeader = make(http.Header)

	emptyAuthHeader.Set("Authorization", "")
	malformedTagAuthHeader.Set("Authorization", "ApiKoy")
	malformedLengthAuthHeader.Set("Authorization", "ApiKey")
	//noKeyAuthHeader.Set("Authorization", "ApiKey ")
	goodKeyAuthHeader.Set("Authorization", "ApiKey g00dk3y")

	testCases := map[string]struct {
		input http.Header
		want  string
		error error
	}{
		"no auth header":               {noAuthHeader, "", errors.New("no authorization header included")},
		"empty auth header":            {emptyAuthHeader, "", errors.New("no authorization header included")},
		"malformed tag auth header":    {malformedTagAuthHeader, "", errors.New("malformed authorization header")},
		"malformed length auth header": {malformedLengthAuthHeader, "", errors.New("malformed authorization header")},
		//"no key auth header":           {noKeyAuthHeader, "", errors.New("no authorization header included")},
		"good key auth header": {goodKeyAuthHeader, "g00dk3y", nil},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.input)
			if !reflect.DeepEqual(tc.want, got) || !reflect.DeepEqual(tc.error, err) {
				t.Fatalf("GetAPIKey (%s) - want: %v, %v got: %v, %v", name, tc.want, tc.error, got, err)
			}
		})
	}

}
