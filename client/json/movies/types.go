package movies

import "project/github/comics/client/json/common"

const Movies common.DataType = "Movies"

type MovieResponseData struct {
	DataType common.DataType `json:"datatype"`
	RawData  *MainType       `json:"rawdata"`
}
