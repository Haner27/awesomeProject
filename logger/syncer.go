package logger

import (
	"fmt"
	"go.uber.org/zap/zapcore"
	"gopkg.in/olivere/elastic.v5"
	"io/ioutil"
	"os"
)

var DevNullSyncEr = zapcore.AddSync(ioutil.Discard)
var StdoutSyncEr = zapcore.Lock(os.Stdout)
var StderrSyncEr = zapcore.Lock(os.Stderr)

func New()  {

}

type ElasticSearchSyncEr struct {
	esClient *elastic.Client
}

func NewElasticSearchSyncEr(esClient *elastic.Client) zapcore.WriteSyncer {
	esSyncer := &ElasticSearchSyncEr{
		esClient: esClient,
	}
	return zapcore.AddSync(esSyncer)
}

func(es *ElasticSearchSyncEr) Write(p []byte) (n int, err error) {
	fmt.Println("Write")
	return 0, nil
}

func(es *ElasticSearchSyncEr) Sync() error {
	fmt.Println("Sync")
	return nil
}
