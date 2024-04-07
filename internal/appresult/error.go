package appresult

import "encoding/json"

var (
	ErrMissingParam           = NewAppError(nil, "missing param", "SE-00001")
	ErrNotFound               = NewAppError(nil, "not found", "SE-00003")
	ErrInternalServer         = NewAppError(nil, "internal server error", "SE-00004")
	ErrNotData                = NewAppError(nil, "data not found", "SE-00005")
	ErrSendingSmsLimitFull    = NewAppError(nil, "sms sending limit full, try after one hour", "SE-00006")
	ErrNotAcceptable          = NewAppError(nil, "Not acceptable", "SE-00007")
	ErrCheckPasswordLimitFull = NewAppError(nil, "ErrCheckPasswordLimitFull", "SE-00008")
	ErrUserIsBlock            = NewAppError(nil, "user blockda", "SE-00009")
	ErrEmptyValue             = NewAppError(nil, "Bad Request: value not found", "SE-00010")
)

type AppError struct {
	Status  bool   `json:"status"`
	Err     error  `json:"-"`
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error { return e.Err }

func (e *AppError) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return marshal
}

func NewAppError(err error, message, code string) *AppError {
	return &AppError{
		Status:  false,
		Err:     err,
		Message: message,
		Code:    code,
	}
}

func systemError(err error) *AppError {
	return NewAppError(err, "internal system error", "SE-000")
}
