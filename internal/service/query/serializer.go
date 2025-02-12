package query

type Request struct {
	Filters Filters `json:"filters"`
	Page    int     `json:"page"`
	Count   int     `json:"count"`
}

type Filters struct {
	Severity string `json:"severity"`
}

type Response struct {
	Pagination Pagination `json:"pagination"`
	Data       []Payload  `json:"data"`
}

type Pagination struct {
	Page    int  `json:"page"`
	Count   int  `json:"count"`
	HasNext bool `json:"has_next"`
}

type Payload struct {
	Id             string   `json:"id"`
	Severity       string   `json:"severity"`
	Cvss           float32  `json:"cvss"`
	Status         string   `json:"status"`
	PackageName    string   `json:"package_name"`
	CurrentVersion string   `json:"current_version"`
	FixedVersion   string   `json:"fixed_version"`
	Description    string   `json:"description"`
	PublishedDate  string   `json:"published_date"`
	Link           string   `json:"link"`
	RiskFactors    []string `json:"risk_factors"`
}
