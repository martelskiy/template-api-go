package api_response

type InternalServerErrorProblemDetails struct {
	*ProblemDetails
}

func NewInternalServerErrorProblemDetails() *InternalServerErrorProblemDetails {
	return &InternalServerErrorProblemDetails{
		&ProblemDetails{
			Status:  500,
			Type:    typeUrlPattern + "500",
			Title:   "Internal Server Error",
			Details: "An unexpected error occured on the server and has been logged",
		},
	}
}
