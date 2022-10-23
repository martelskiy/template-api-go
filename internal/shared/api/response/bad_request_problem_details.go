package api_response

type BadRequestProblemDetails struct {
	*ProblemDetails
	Errors []ProblemDetailsError
}

func NewBadRequestProblemDetails(er []ProblemDetailsError) *BadRequestProblemDetails {
	return &BadRequestProblemDetails{
		ProblemDetails: &ProblemDetails{
			Status:  400,
			Type:    typeUrlPattern + "400",
			Title:   "Bad Request",
			Details: "The request produced one or more errors",
		},
		Errors: er,
	}
}
