package people

import "project/github/comics/client/json/common"

const Movie common.DataType = "Movie"

type MovieResponseData struct {
	DataType common.DataType `json:"datatype"`
	RawData  *MainType       `json:"rawdata"`
}
