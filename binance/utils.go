package binance

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
)

func generateSignature(queryString, body, secret string) string {
	totalParams := queryString + body

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(totalParams))

	return hex.EncodeToString(h.Sum(nil))
}

func formQueryString() {

}

func parseBody(res *http.Response, to any) error {
	body, readError := io.ReadAll(res.Body)
	if readError != nil {
		return readError
	}

	if unmarshalError := json.Unmarshal(body, to); unmarshalError != nil {
		return unmarshalError
	}

	return nil
}
