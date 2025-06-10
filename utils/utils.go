package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofrs/uuid"
	//"github.com/go-playground/validator/v10"
)

/*var Validate = validator.New()

func Validator(v any, w http.ResponseWriter) error {
	if err := Validate.Struct(v); err != nil {
		errors := err.(validator.ValidationErrors)
		WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return errors
	}
	return nil
}*/

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

func ParseJSON(r *http.Request, v any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(v)
}

func GetParam(r *http.Request, name string) (string, error) {
	param := r.PathValue(name)
	if param != "" {
		return param, nil
	}
	return "", fmt.Errorf("missing parameter: %s", name)
}

func GetParamInt(r *http.Request, name string) (int, error) {
	param, err := GetParam(r, name)
	if err != nil {
		return 0, err
	}

	num, err := strconv.Atoi(param)
	if err != nil {
		return 0, fmt.Errorf("invalid parameter %s: %v", name, err)
	}
	return num, nil

}

func GetParamUUID(r *http.Request, name string) (uuid.UUID, error) {
	param, err := GetParam(r, name)
	if err != nil {
		return uuid.Nil, err
	}

	id, err := uuid.FromString(param)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid UUID parameter %s: %v", name, err)
	}
	return id, nil
}

func StringtoUUID(s string) (uuid.UUID, error) {
	id, err := uuid.FromString(s)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid UUID string: %v", err)
	}
	return id, nil
}

func GetTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	tokenQuery := r.URL.Query().Get("token")

	if tokenAuth != "" {
		var tokenValue string
		fmt.Sscanf(tokenAuth, "Bearer %s", &tokenValue)
		if tokenValue == "" {
			return ""
		}
		return tokenValue
	}

	if tokenQuery != "" {
		return tokenQuery
	}

	tokenCookie, err := r.Cookie("token")
	if err == nil {
		return tokenCookie.Value
	}

	return ""
}
