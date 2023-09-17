package contract

import (
	"project/github/comics/client/json/common"
)

type InformationCaller[T any] interface {
	GetData() (T, error)
	GetDataType() common.DataType
}
