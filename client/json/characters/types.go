package characters

import "project/github/comics/client/json/common"

const Characters common.DataType = "Characters"

type CharactersArray []MainType

type CharactersResponseData struct {
	DataType common.DataType `json:"datatype"`
	RawData  *MainType       `json:"rawdata"`
}
