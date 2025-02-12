package scan

type ScanRequest struct {
	Repo  string   `json:"repo"`
	Files []string `json:"files"`
}

type ScanResponse struct {
	Message string `json:"msg"`
}
