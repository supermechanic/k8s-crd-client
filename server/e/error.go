package e

const (
	OK                int32 = 0
	ERROR_BAD_PARAM         = 1001
	ERROR_INTERNAL          = 1002
	ERROR_NOT_SUPPORT       = 1003
	ERROR_DB_NOT_FIND       = 1004
)

var message = map[int32]string{
	OK:                "success",
	ERROR_BAD_PARAM:   "bad request parameter",
	ERROR_INTERNAL:    "internal error",
	ERROR_NOT_SUPPORT: "not supported choice",
	ERROR_DB_NOT_FIND: "con't find request data",
}

//GetMessage ///
func GetMessage(code int32) string {
	return message[code]
}
