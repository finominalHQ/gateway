package route

type AuthType string

var (
	NO_AUTH    AuthType = "NO_AUTH"
	BASIC_AUTH AuthType = "BASIC_AUTH"
	TOKEN_AUTH AuthType = "TOKEN_AUTH"
	PIN_AUTH   AuthType = "PIN_AUTH"
	OTP_AUTH   AuthType = "OTP_AUTH"
)

func (s AuthType) String() string {
	return string(s)
}

type StatusType string

const (
	ACTIVE   StatusType = "ACTIVE"
	INACTIVE StatusType = "INACTIVE"
)

func (s StatusType) String() string {
	return string(s)
}

type TypeType string

const (
	INCOMING TypeType = "INCOMING"
	OUTGOING TypeType = "OUTGOING"
)

func (s TypeType) String() string {
	return string(s)
}
