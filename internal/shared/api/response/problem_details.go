package api_response

type ProblemDetails struct {
	Status  int    `json:"status"`
	Type    string `json:"type"`
	Title   string `json:"title"`
	Details string `json:"details"`
}

type ProblemDetailsError struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

const typeUrlPattern string = "https://httpstatuses/"
