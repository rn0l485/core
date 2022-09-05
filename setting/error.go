package setting 


import (
	"errors"
)


var (
	Err_Internal = errors.New("internal-error")
	Err_Payload_format = errors.New("payload-format-error")
	Err_Unauthorized = errors.New("unauthorized")
	
	Err_DataDuplicated = errors.New("data-duplicated")
	Err_IsExpired = errors.New("is-expired")
	Err_NoData = errors.New("no-data")
	Err_DataExist = errors.New("data-exist")

	Err_NotSupport = errors.New("not-support")
	Err_Uncompleted = errors.New("uncompleted")
)
