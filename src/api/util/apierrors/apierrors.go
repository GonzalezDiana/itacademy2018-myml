package apierrors

type ApiError struct {
    StatusCode int    `json:"status_code"`
    Error      string `json:"error"`
}