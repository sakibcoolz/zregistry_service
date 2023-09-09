package literals

const (
	REQUEST_VALIDATION_ERROR = "REQUEST_VALIDATION_ERROR"
	FAIL_TO_STORE            = "FAIL_TO_STORE"
	LOGIN_FAILED             = "LOGIN_FAILED"
)

// var (
// 	ErrorResponse = make(map[string]string)
// )

var ErrorResponse = map[string]string{
	"REQUEST_VALIDATION_ERROR": "Failed Validate Body",
	"FAIL_TO_STORE":            "Failed to store please try again after some time",
	"LOGIN_FAILED":             "Incorrect username and password",
}
