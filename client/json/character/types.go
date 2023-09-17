package character

import "project/github/comics/client/json/common"

const Character common.DataType = "Character"

type CharacterResponseData struct {
	DataType common.DataType `json:"datatype"`
	RawData  *MainType       `json:"rawdata"`
}
