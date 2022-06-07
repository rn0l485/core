package setting 


import (
	"errors"
)


var (
	Err_Internal = errors.New("internal-error")
	Err_Payload_format = errors.New("payload-format-error")
	Err_Unauthorized = errors.New("unauthorized")
	
	Err_DataDuplicated = errors.New("data-duplicated")
	Err_NoData = errors.New("no-data")
)
