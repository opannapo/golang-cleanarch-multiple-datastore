package constant

const (
	//ErrDuplicateCredentialUserName --> Username already exists
	ErrDuplicateCredentialUserName = 1
)

var errorMessage = map[int]string{
	ErrDuplicateCredentialUserName: "Username Tidak Tersedia!",
}

//GetErrorMessage Parsing from int code to String as Error Message
func GetErrorMessage(code int) string {
	return errorMessage[code]
}
