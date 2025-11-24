package utils

import "strings"

type APIResponse struct {
	Success bool   `json:"success,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func IsEmpty(key string) bool {
	return strings.TrimSpace(key) == ""
}
