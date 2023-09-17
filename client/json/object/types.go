package object

import "project/github/comics/client/json/common"

const Object common.DataType = "Object"

type ObjectResponseData struct {
	DataType common.DataType `json:"datatype"`
	RawData  *MainType       `json:"rawdata"`
}
