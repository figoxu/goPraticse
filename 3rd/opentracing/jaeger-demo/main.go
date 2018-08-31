package main

import (
	"github.com/uber/jaeger-lib/metrics"
	"log"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"net/http"
	"time"
)

func main() {
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}
	// Example logger and metrics factory. Use github.com/uber/jaeger-client-go/log
	// and github.com/uber/jaeger-lib/metrics respectively to bind to real logging and metrics
	// frameworks.
	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	// Initialize tracer with a logger and a metrics factory
	closer, err := cfg.InitGlobalTracer(
		"test-service",
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	if err != nil {
		log.Printf("Could not initialize jaeger tracer: %s", err.Error())
		return
	}
	defer closer.Close()

	engin := gin.Default()
	engin.Use(jaegerMiddleware)
	engin.GET("/test", testFunc)
	engin.GET("/hello", helloFunc)
	engin.POST("/postFunc", postFunc)

	engin.Run(":8080")
}

func jaegerMiddleware(c *gin.Context) {
	span := opentracing.GlobalTracer().StartSpan(c.Request.RequestURI)
	defer span.Finish()
	c.Set("mainspan", span)
	span.SetTag("spanContext", span.Context())
	span.SetTag("functionName", c.HandlerName())
	span.SetTag("Method", c.Request.Method)
	span.SetTag("Params", c.Params)

	c.Next()
}

func testFunc(c *gin.Context) {
	span, err := getSpan(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, 500)
		return
	}
	doSomthing(span, "wo shi fang fa can shu")

	c.JSON(http.StatusOK, "testFunc 200")
}

func doSomthing(span opentracing.Span, params string) {
	child := opentracing.GlobalTracer().StartSpan(
		"main.doSomthing",
		opentracing.ChildOf(span.Context()),
	)
	child.SetTag("Params", params)

	child.SetTag("result", "success")

	defer child.Finish()
	time.Sleep(1 * time.Second)
}

func postFunc(c *gin.Context) {
	span, err := getSpan(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, 500)
		return
	}
	doSomthing(span, "canshu")

	c.JSON(http.StatusOK, "postFunc 200")
}

func helloFunc(c *gin.Context) {
	//parentSpan, err := c.Get("mainspan")
	//if !err {
	//	c.JSON(http.StatusBadRequest, 400)
	//}
	//if span, ok := (parentSpan).(opentracing.Span); ok {
	//	child := opentracing.GlobalTracer().StartSpan(
	//		c.Request.RequestURI, opentracing.ChildOf(span.Context()))
	//	defer child.Finish()
	//}
	c.JSON(http.StatusOK, "helloFunc 200")
}

func getSpan(c *gin.Context) (opentracing.Span, error) {
	parentSpan, err := c.Get("mainspan")
	if !err {
		return nil, errors.New("get main span error")
	}

	if span, ok := (parentSpan).(opentracing.Span); ok {
		return span, nil
	}
	return nil, errors.New("get main span error")
}
