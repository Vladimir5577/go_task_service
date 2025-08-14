package model

type ServiceError struct {
	StatusCode int
	Message    string
}

func (e *ServiceError) Error() string {
	return e.Message
}

func (e *ServiceError) GetStatusCode() int {
	return e.StatusCode
}
