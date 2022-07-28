package route

type AuthType string

var (
	NO_AUTH    = "NO_AUTH"
	BASIC_AUTH = "BASIC_AUTH"
	TOKEN_AUTH = "TOKEN_AUTH"
	PIN_AUTH   = "PIN_AUTH"
	OTP_AUTH   = "OTP_AUTH"
)

type StatusType string

const (
	ACTIVE   = "ACTIVE"
	INACTIVE = "INACTIVE"
)

type TypeType string

const (
	INCOMING = "INCOMING"
	OUTGOING = "OUTGOING"
)
