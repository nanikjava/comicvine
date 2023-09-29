package client

import (
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/nanikjava/comicstype/json/common"
	"log"
	"time"
)

func New(apikey string) *req.Client {
	client := req.C().
		SetCommonQueryParam("api_key", apikey).
		SetTimeout(5 * time.Second)

	return client
}

func Call(c common.CommonStruct, queryMap map[string]string, resultType interface{}, urlPath string) interface{} {
	resp, err := c.Client.R().
		SetQueryParams(queryMap).
		SetSuccessResult(resultType).
		EnableDump().
		Get(fmt.Sprintf("%s%s", common.COMICVIEW_BASEURL, urlPath))

	if err != nil {
		log.Printf("err: ", err)
		return nil
	}

	return resp.Request.Result
}

func CallSingle(c common.CommonStruct, queryMap map[string]string, resultType interface{}, urlPath string) interface{} {
	resp, err := c.Client.R().
		SetQueryParams(queryMap).
		SetSuccessResult(resultType).
		EnableDump().
		Get(fmt.Sprintf("%s", urlPath))

	if err != nil {
		log.Printf("err: ", err)
		return nil
	}

	return resp.Request.Result
}
