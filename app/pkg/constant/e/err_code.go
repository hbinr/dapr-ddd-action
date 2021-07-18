package e

// ResCode .
type ResCode int

const (
	CodeSuccess       ResCode = 200
	CodeInvalidParams ResCode = 400
	CodeError         ResCode = 500

	CodeConvDataErr       ResCode = 500000
	CodeValidateParamsErr ResCode = 500001
	CodeInvalidToken      ResCode = 500002
	CodeNeedLogin         ResCode = 500003
	CodeInvalidID         ResCode = 500004

	CodeWrongPassword           ResCode = 403001
	CodeWrongUserNameOrPassword ResCode = 403002
	CodeUserNotExist            ResCode = 404001
	CodeUserExist               ResCode = 409002
	CodeEmailExist              ResCode = 409003
)
