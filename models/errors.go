package models

type DoesNotExist struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ValidationError struct {
	ValidationErrors map[string]string `json:"validation_errors"`
}
