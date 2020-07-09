package logger

import (
	"awesomeProject/mq/kafka"
	"context"
	"fmt"
	"go.uber.org/zap/zapcore"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/olivere/elastic.v5"
	"io/ioutil"
	"os"
	"time"
)

var DevNullSyncEr = zapcore.AddSync(ioutil.Discard)
var StdoutSyncEr = zapcore.Lock(os.Stdout)
var StderrSyncEr = zapcore.Lock(os.Stderr)

// es syncer
type ElasticSearchSyncEr struct {
	esClient *elastic.Client
}

func NewElasticSearchSyncEr(esClient *elastic.Client) zapcore.WriteSyncer {
	esSyncer := &ElasticSearchSyncEr{
		esClient: esClient,
	}
	return zapcore.AddSync(esSyncer)
}

func (es *ElasticSearchSyncEr) Write(p []byte) (n int, err error) {
	logIndexName := fmt.Sprintf("log-%s", time.Now().Format("20060102"))
	logID := bson.NewObjectId().Hex()
	_, err = es.esClient.Index().Index(logIndexName).Type("main").Id(logID).BodyString(string(p)).Do(context.Background())
	return 0, err
}

// kafka syncer
type KafkaSyncEr struct {
	Producer *kafka.KafkaProducer
}

func NewKafkaSyncEr(producer *kafka.KafkaProducer) zapcore.WriteSyncer {
	kfkSyncer := &KafkaSyncEr{
		Producer: producer,
	}
	return zapcore.AddSync(kfkSyncer)
}

func (kfk *KafkaSyncEr) Write(p []byte) (n int, err error) {
	logID := bson.NewObjectId().Hex()
	kfk.Producer.SendMsg(logID, string(p))
	return 0, err
}