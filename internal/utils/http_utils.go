package utils

// JSONError is a struct for error response
// @Description JSONResult is a generic JSON response
type JSONResult struct {
	Success bool        `json:"success" example:"true"`             // Success is a boolean value indicating whether the request was successful or not
	Message string      `json:"message" example:"Response message"` // Message is a string value containing a message about the request
	Error   interface{} `json:"error,omitempty"`                    // Error is an JSONError containing information about the error
	Data    interface{} `json:"data,omitempty"`                     // Data is an interface{} value containing the data returned by the request
} // @name JSONResult

// JSONError is a struct for error response
// @Description JSONError is a generic JSON error response
type JSONError struct {
	Code    int    `json:"code" example:"400"`              // Code is an integer value containing the HTTP status code of the error
	Message string `json:"message" example:"Error message"` // Message is a string value containing the error message
} // @name JSONError
