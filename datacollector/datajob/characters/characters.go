package characters

import (
	"errors"
	"fmt"
	"github.com/nanikjava/comicstype/contract"
	client "github.com/nanikjava/comicstype/http"
	"github.com/nanikjava/comicstype/json/characters"
	"github.com/nanikjava/comicstype/json/common"
	"project/github/comics/async"

	"strconv"
)

type Characters struct {
	common.CommonStruct
	arr characters.CharactersArray
}

func (c Characters) Len() int {
	return len(c.arr)
}
func (c Characters) Get(idx int) *characters.MainType {
	return &c.arr[idx]
}

func (c Characters) GetDataType() common.DataType {
	return characters.Characters
}

func (c *Characters) GetData(urlString string) error {
	var dataType *characters.MainType
	var totalRecords = 3

	var err error
	var cType *characters.MainType

	for i := 0; i < totalRecords; i++ {
		mapQuery := map[string]string{
			"format": "json",
			"limit":  strconv.Itoa(c.Limit),
			"offset": strconv.Itoa(c.Offset),
		}

		cType, err = getData(*c, mapQuery, dataType)

		if cType != nil {
			for _, result := range cType.Results {
				responseData := &characters.CharactersResponseData{
					DataType: characters.Characters,
					Data:     &result,
				}
				async.SendToMQ(responseData, "characters-exchange")
			}

			c.Offset = c.Offset + cType.NumberOfPageResults
			fmt.Println("Next offset = ", c.Offset)

			c.arr = append(c.arr, *cType)
		}
	}
	return err
}

func getData(c Characters, queryMap map[string]string, resultType *characters.MainType) (*characters.MainType, error) {
	resp := c.Call(queryMap, resultType, "characters")

	if resp == nil {
		return nil, errors.New("Error getting data")
	}

	k, ok := resp.(*characters.MainType)
	if ok {
		return k, nil
	}
	return nil, nil
}

func New(apikey string) contract.InformationCaller[characters.CharactersArray, characters.MainType] {
	c := &Characters{
		CommonStruct: common.CommonStruct{
			Offset: 0,
			Limit:  10,
		},
		arr: []characters.MainType{},
	}
	c.HttpClient = client.New(apikey)
	return c
}
