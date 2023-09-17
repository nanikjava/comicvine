package client

import (
	"github.com/imroc/req/v3"
	"time"
)

func CreateClient(apikey string) *req.Client {
	client := req.C().
		SetCommonQueryParam("api_key", apikey).
		SetTimeout(5 * time.Second)

	return client
}
