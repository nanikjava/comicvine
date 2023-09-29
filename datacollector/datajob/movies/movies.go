package movies

import (
	"errors"
	"fmt"
	"github.com/nanikjava/comicstype/contract"
	"github.com/nanikjava/comicstype/json/common"
	"github.com/nanikjava/comicstype/json/movies"
	"project/github/comics/client"
	"project/github/comics/client/async"
	http "project/github/comics/client/sync"
	"strconv"
)

type Movies struct {
	common.CommonStruct
}

func (c Movies) Get(idx int) *movies.MainType {
	//TODO implement me
	panic("implement me")
}

func (c Movies) Len() int {
	//TODO implement me
	panic("implement me")
}

func (c Movies) GetDataType() common.DataType {
	return movies.Movies
}

func (c Movies) GetData(urlString string) error {
	var dataType *movies.MainType
	var totalRecords = 3

	var err error
	var cType *movies.MainType

	for i := 0; i < totalRecords; i++ {
		mapQuery := map[string]string{
			"format": "json",
			"limit":  strconv.Itoa(c.Limit),
			"offset": strconv.Itoa(c.Offset),
		}

		cType, err = getData(c, mapQuery, dataType)

		if cType != nil {
			responseData := &movies.MovieResponseData{
				DataType: movies.Movies,
				RawData:  cType,
			}
			async.SendToMQ(responseData, "movies")
			c.Offset = c.Offset + cType.NumberOfPageResults
			fmt.Println("Next offset = ", c.Offset)
		}
	}
	return err
}

func getData(c Movies, queryMap map[string]string, resultType *movies.MainType) (*movies.MainType, error) {
	resp := http.Call(c.CommonStruct, queryMap, resultType, "characters")

	if resp == nil {
		return nil, errors.New("Error getting data")
	}

	k, ok := resp.(*movies.MainType)
	if ok {
		return k, nil
	}
	return nil, nil
}

func New(apikey string) contract.InformationCaller[movies.MainType, movies.MainType] {
	c := &Movies{
		CommonStruct: common.CommonStruct{
			Offset: 0,
			Limit:  10,
		},
	}
	c.Client = client.CreateClient(apikey)
	return c
}
