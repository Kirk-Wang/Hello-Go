package main

import (
	"testing"
	"time"

	config "github.com/Kirk-Wang/Hello-Gopher/history/17.9/crawler/config"
	config2 "github.com/Kirk-Wang/Hello-Gopher/history/17.9/crawler_distributed/config"

	"github.com/Kirk-Wang/Hello-Gopher/history/17.9/crawler_distributed/rpcsupport"
	"github.com/Kirk-Wang/Hello-Gopher/history/17.9/crawler_distributed/worker"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "http://album.zhenai.com/u/1113093123",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: `{"Name":"心想事成","Gender":"女士"}`,
		},
	}

	var result worker.ParseResult
	err = client.Call(config2.CrawlServiceRpc, req, &result)

	if err != nil {
		t.Error(err)
	} else {
		// t.Errorf("%v", result)
		// fmt.Println(result)
	}
}
