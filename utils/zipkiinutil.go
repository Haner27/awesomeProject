package utils

import (
	"github.com/opentracing/opentracing-go"
	zipKinOpenTracing "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/reporter"
	zipKinHttp "github.com/openzipkin/zipkin-go/reporter/http"
	"log"
)

func InitTracer(zipKinHostPort, serverName, serverHostPort string) reporter.Reporter {
	// set up a span reporter
	reporter := zipKinHttp.NewReporter("http://" + zipKinHostPort + "/api/v2/spans")
	// create our local service endpoint
	endpoint, err := zipkin.NewEndpoint(serverName, serverHostPort)
	if err != nil {
		log.Fatalf("unable to create local endpoint: %+v\n", err)
	}
	// initialize our tracer
	nativeTracer, err := zipkin.NewTracer(
		reporter,
		zipkin.WithLocalEndpoint(endpoint),
		zipkin.WithSharedSpans(false),
	)
	if err != nil {
		log.Fatalf("unable to create tracer: %+v\n", err)
	}
	// use zipkin-go-opentracing to wrap our tracer
	tracer := zipKinOpenTracing.Wrap(nativeTracer)
	// optionally set as Global OpenTracing tracer instance
	opentracing.SetGlobalTracer(tracer)
	return reporter
}
