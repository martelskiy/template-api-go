package api_response

type NotFoundProblemDetails struct {
	*ProblemDetails
}

func NewNotFoundProblemDetails() *InternalServerErrorProblemDetails {
	return &InternalServerErrorProblemDetails{
		ProblemDetails: &ProblemDetails{
			Status:  404,
			Type:    typeUrlPattern + "404",
			Title:   "Not Found",
			Details: "The requested resource could not be found",
		},
	}
}
