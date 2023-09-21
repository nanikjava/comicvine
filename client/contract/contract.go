package contract

import (
	"project/github/comics/client/json/common"
)

type InformationCaller[T any, V any] interface {
	GetData(apiUrl string) error
	GetDataType() common.DataType
	Get(idx int) *V
	Len() int
}
