package errorx

type ErrorType struct {
	t string
}

var (
	ErrorTypeUnknown        = ErrorType{"unknown"}
	ErrorTypeAuthorization  = ErrorType{"authorization"}
	ErrorTypeIncorrectInput = ErrorType{"incorrect-input"}
	ErrorTypeConvertData    = ErrorType{"convert-data"}
)

type SlugErr struct {
	error     string    // 错误信息, 展示给用户
	slug      string    // slug to  URL
	errorType ErrorType // error check
}

func (s SlugErr) Error() string {
	return s.error
}

func (s SlugErr) Slug() string {
	return s.slug
}

func (s SlugErr) ErrorType() ErrorType {
	return s.errorType
}

func NewSlugError(error string, slug string) SlugErr {
	return SlugErr{
		error:     error,
		slug:      slug,
		errorType: ErrorTypeUnknown,
	}
}

func NewAuthorizationError(error string, slug string) SlugErr {
	return SlugErr{
		error:     error,
		slug:      slug,
		errorType: ErrorTypeAuthorization,
	}
}

func NewIncorrectInputError(error string, slug string) SlugErr {
	return SlugErr{
		error:     error,
		slug:      slug,
		errorType: ErrorTypeIncorrectInput,
	}
}

func NewConvertDataError(error string, slug string) SlugErr {
	return SlugErr{
		error:     error,
		slug:      slug,
		errorType: ErrorTypeConvertData,
	}
}
