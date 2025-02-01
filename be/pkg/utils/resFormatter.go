package utils

type ResFormatter struct {
	Success    bool   `json:"success"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}
