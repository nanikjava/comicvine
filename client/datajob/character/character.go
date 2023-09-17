package character

import (
	"errors"
	"fmt"
	"project/github/comics/client"
	"project/github/comics/client/async"
	"project/github/comics/client/contract"
	character "project/github/comics/client/json/character"
	"project/github/comics/client/json/common"
	http "project/github/comics/client/sync"

	"strconv"
)

type Character struct {
	common.CommonStruct
}

func (c Character) GetDataType() common.DataType {
	return character.Character
}

func (c Character) GetData() (*character.MainType, error) {
	var dataType *character.MainType
	var totalRecords = 3

	var err error
	var cType *character.MainType

	for i := 0; i < totalRecords; i++ {
		mapQuery := map[string]string{
			"format": "json",
			"limit":  strconv.Itoa(c.Limit),
			"offset": strconv.Itoa(c.Offset),
		}

		cType, err = getData(c, mapQuery, dataType)

		if cType != nil {
			async.SendToMQ(character.Character, cType)
			c.Offset = c.Offset + cType.NumberOfPageResults
			fmt.Println("Next offset = ", c.Offset)
		}
	}
	return cType, err
}

func getData(c Character, queryMap map[string]string, resultType *character.MainType) (*character.MainType, error) {
	resp := http.Call(c.CommonStruct, queryMap, resultType, "characters")

	if resp == nil {
		return nil, errors.New("Error getting data")
	}

	k, ok := resp.(*character.MainType)
	if ok {
		return k, nil
	}
	return nil, nil
}

func New(apikey string) contract.InformationCaller[*character.MainType] {
	c := Character{
		CommonStruct: common.CommonStruct{
			Offset: 0,
			Limit:  10,
		},
	}
	c.Client = client.CreateClient(apikey)
	return c
}
