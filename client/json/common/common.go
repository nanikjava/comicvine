package common

import "github.com/imroc/req/v3"

type DataType string

type CommonStruct struct {
	Offset int
	Limit  int
	Client *req.Client
}

const (
	COMICVIEW_BASEURL = "https://comicvine.gamespot.com/api/"
)
