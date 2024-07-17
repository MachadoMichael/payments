package encrypt

import (
	"encoding/json"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashData(data interface{}) (string, error) {
	newData, err := json.Marshal(data)
	hashedData, err := bcrypt.GenerateFromPassword([]byte(newData), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("could not hash password %w", err)
	}

	return string(hashedData), nil
}
