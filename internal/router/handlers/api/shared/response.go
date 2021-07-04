package shared

type SuccessErrResp struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}
