package es

import (
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"strings"
	"time"
)
var EsClient *elastic.Client

func InitEsDao(hosts []string) {
	var refinedHosts []string
	for _, host := range hosts {
		if !strings.HasPrefix(host, "http://") || !strings.HasPrefix(host, "https://") {
			host = fmt.Sprintf("http://%s", host)
		}
		refinedHosts = append(refinedHosts, host)
	}
	EsClient = connect(refinedHosts)
}

func connect(hosts []string) *elastic.Client {
	client, err := elastic.NewClient(
		elastic.SetURL(hosts...),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetGzip(true),
	)

	if err != nil {
		panic(err)
	}
	return client
}