package shared

type SuccessErrResp struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}
type SuccessIDErrResp struct {
	ID      int64  `json:"id,omitempty"`
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}
