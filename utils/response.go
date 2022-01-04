package utils

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(data interface{}) Response {
	return Response{"success", data}
}

func Error(message string, data interface{}) Response {
	return Response{message, data}
}
