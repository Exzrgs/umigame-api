package myerrors

type ErrCode string

const (
	Unknown ErrCode = "U000"

	InsertDataFailed ErrCode = "D000"
	GetDataFailed    ErrCode = "D001"
	UpdateDataFailed ErrCode = "D002"
	DeleteDataFailed ErrCode = "D003"

	NoData           ErrCode = "S000"
	EncryptFailed    ErrCode = "S001"
	GenUUIDFailed    ErrCode = "S002"
	SendMailFailed   ErrCode = "S003"
	TypeCastFailed   ErrCode = "S004"
	SQLPrepareFailed ErrCode = "S005"

	ReqDecodeFailed ErrCode = "R000"
	BadParameter    ErrCode = "R001"
	EmailInvalid    ErrCode = "R002"
	BadCookie       ErrCode = "R002"

	InvalidPassword ErrCode = "U000"
	NotActivate     ErrCode = "U001"
	InvalidUUID     ErrCode = "U002"
)
