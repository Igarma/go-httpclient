package gohttpclient

import (
	"github.com/Igarma/go-httpclient/gohttp"
)

func exampleUsage(){
	Client := gohttp.HttpClient{}

	Client.Get()
}