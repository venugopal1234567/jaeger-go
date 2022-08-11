package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

const (
	serviceName      = "my-first-tracer"
	environment      = "dev"
	id               = 1
	jaeger_host_port = "jaeger-agent:6831"
)

func main() {

	cfg := &config.Configuration{
		ServiceName: serviceName,

		// "const" sampler is a binary sampling strategy: 0=never sample, 1=always sample.
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},

		// Log the emitted spans to stdout.
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: jaeger_host_port,
		},
	}

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	defer closer.Close()

	opentracing.SetGlobalTracer(tracer)

	http.HandleFunc("/helloworld", func(w http.ResponseWriter, r *http.Request) {

		spanCtx, _ := tracer.Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(r.Header),
		)
		span := tracer.StartSpan("helloworld: GET", ext.RPCServerOption(spanCtx))
		defer span.Finish()
		fmt.Fprintf(w, "Hello, World!")
	})
	fmt.Printf("Server running (port=8080), route: http://localhost:8080/helloworld\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
