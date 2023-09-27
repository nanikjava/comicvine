package characters

import "project/github/comics/client/json/common"

const Characters common.DataType = "Characters"

type CharactersArray []MainType

type CharactersResponseData struct {
	DataType common.DataType `json:"datatype"`
	Data     *Results        `json:"data"`
}
