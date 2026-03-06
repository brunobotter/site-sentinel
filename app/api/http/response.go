package http

import _http "net/http"

type HttpResponse struct {
	StatusCode   int
	Body         any
	ErrorMessage string
}

type HttpResponseData struct {
	Data any `json:"data"`
}

func Ok(body any) *HttpResponse {
	return &HttpResponse{
		StatusCode: _http.StatusOK,
		Body:       HttpResponseData{body},
	}
}

func OkNoContent() *HttpResponse {
	return &HttpResponse{
		StatusCode: _http.StatusNoContent,
	}
}

func Created(bodyData any) *HttpResponse {
	return &HttpResponse{
		StatusCode: _http.StatusCreated,
		Body:       HttpResponseData{bodyData},
	}
}

func CreatedNoContent() *HttpResponse {
	return &HttpResponse{
		StatusCode: _http.StatusCreated,
	}
}

func Unauthorized(message string) *HttpResponse {
	return &HttpResponse{
		StatusCode:   _http.StatusUnauthorized,
		Body:         buildErrorBody(message),
		ErrorMessage: message,
	}
}

func Forbidden(message string) *HttpResponse {
	return &HttpResponse{
		StatusCode:   _http.StatusForbidden,
		Body:         buildErrorBody(message),
		ErrorMessage: message,
	}
}

func UnprocessableEntity(message string) *HttpResponse {
	return &HttpResponse{
		StatusCode:   _http.StatusUnprocessableEntity,
		Body:         buildErrorBody(message),
		ErrorMessage: message,
	}
}

func NotFound(message string) *HttpResponse {
	return &HttpResponse{
		StatusCode:   _http.StatusNotFound,
		Body:         buildErrorBody(message),
		ErrorMessage: message,
	}
}

func ServiceUnavailable(message string) *HttpResponse {
	return &HttpResponse{
		StatusCode:   _http.StatusServiceUnavailable,
		Body:         buildErrorBody(message),
		ErrorMessage: message,
	}
}

func InternalServerError(message string) *HttpResponse {
	return &HttpResponse{
		StatusCode:   _http.StatusInternalServerError,
		Body:         buildErrorBody(message),
		ErrorMessage: message,
	}
}

func TimeoutExceeded(message string) *HttpResponse {
	return &HttpResponse{
		StatusCode:   _http.StatusGatewayTimeout,
		Body:         buildErrorBody(message),
		ErrorMessage: message,
	}
}

func BadRequest(message string) *HttpResponse {
	return &HttpResponse{
		StatusCode:   _http.StatusBadRequest,
		Body:         buildErrorBody(message),
		ErrorMessage: message,
	}
}

func BadGateway(message string) *HttpResponse {
	return &HttpResponse{
		StatusCode:   _http.StatusBadGateway,
		Body:         buildErrorBody(message),
		ErrorMessage: message,
	}
}

func buildErrorBody(message string) map[string]any {
	return map[string]any{
		"error":   true,
		"message": message,
	}
}
