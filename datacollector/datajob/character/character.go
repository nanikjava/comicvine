package character

import (
	"errors"
	"github.com/nanikjava/comicstype/contract"
	"github.com/nanikjava/comicstype/json/character"
	"github.com/nanikjava/comicstype/json/common"
	"project/github/comics/async"
	"project/github/comics/client"
	"strconv"
)

type Character struct {
	common.CommonStruct
	arr character.CharacterArray
}

func (c Character) Len() int {
	return len(c.arr)
}

func (c Character) GetDataType() common.DataType {
	return character.Character
}
func (c Character) Get(idx int) *character.MainType {
	return c.arr[idx]
}

func (c *Character) GetData(apiUrl string) error {
	var dataType *character.MainType

	var err error
	var cType *character.MainType

	mapQuery := map[string]string{
		"format": "json",
		"limit":  strconv.Itoa(c.Limit),
	}

	cType, err = getData(*c, mapQuery, dataType, apiUrl)

	if cType != nil {
		responseData := &character.CharacterResponseData{
			DataType: character.Character,
			Data:     &cType.Results,
		}
		async.SendToMQ(responseData, "character-exchange")
		c.arr[0] = cType
	}
	return err
}

func getData(c Character, queryMap map[string]string, resultType *character.MainType, url string) (*character.MainType, error) {
	resp := client.CallSingle(c.CommonStruct, queryMap, resultType, url)

	if resp == nil {
		return nil, errors.New("Error getting data")
	}

	k, ok := resp.(*character.MainType)
	if ok {
		return k, nil
	}
	return nil, nil
}

func New(apikey string) contract.InformationCaller[character.MainType, character.MainType] {
	c := &Character{
		CommonStruct: common.CommonStruct{
			Offset: 0,
			Limit:  1,
			Client: client.New(apikey),
		},
		arr: make(character.CharacterArray, 1),
	}
	return c
}
