package scan

type Request struct {
	Repo  string   `json:"repo"`
	Files []string `json:"files"`
}

type Response struct {
	Message string `json:"msg"`
}
