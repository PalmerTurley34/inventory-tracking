package backend

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func GetApiKey(header http.Header) (string, error) {
	apiKey := header.Get("Authorization")
	if apiKey == "" {
		return "", errors.New("authorization header is not found")
	}
	apiKey, found := strings.CutPrefix(apiKey, "ApiKey ")
	if !found {
		return "", errors.New("malformed auth header, expecting \"ApiKey <key>\"")
	}
	return apiKey, nil
}

func GetUrlUUID(r *http.Request) (uuid.UUID, error) {
	idStr := chi.URLParam(r, "ID")
	if idStr == "" {
		return uuid.UUID{}, fmt.Errorf("did not find {ID} param in url")
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("{ID} is not a valid uuid")
	}
	return id, nil
}
