package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"regexp"
	"strings"
)

func GenerateRandomBytes(size int) (string, error) {
	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func HashWithSHA256(randomStr string) string {
	hash := sha256.Sum256([]byte(randomStr))
	return hex.EncodeToString(hash[:])
}

func MapToJsonString(data map[string]string) (string, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), err
}

func JsonStringToMap(jsonStr string) (map[string]string, error) {
	var result map[string]string
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func ToSlug(input string) string {
	slug := strings.ToLower(input)
	reg := regexp.MustCompile(`[^a-z0-9\s-]`)
	slug = reg.ReplaceAllString(slug, "")
	reg = regexp.MustCompile(`[\s\-_]+`)
	slug = reg.ReplaceAllString(slug, "-")
	return strings.Trim(slug, "-")
}
