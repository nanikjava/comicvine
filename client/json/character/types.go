package character

import "project/github/comics/client/json/common"

const Character common.DataType = "Character"

type CharacterArray []*MainType

type CharacterResponseData struct {
	DataType common.DataType `json:"datatype"`
	Data     *Results        `json:"data"`
}
