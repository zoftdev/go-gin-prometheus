# go-gin-prometheus
[![](https://godoc.org/github.com/zoftdev/go-gin-prometheus?status.svg)](https://godoc.org/github.com/zoftdev/go-gin-prometheus) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Gin Web Framework Prometheus metrics exporter

## Installation

`$ go get github.com/zoftdev/go-gin-prometheus`

## Usage

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zoftdev/go-gin-prometheus"
)

func main() {
	r := gin.New()

	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello world!")
	})

	r.Run(":29090")
}
```

See the [example.go file](https://github.com/zoftdev/go-gin-prometheus/blob/master/example/example.go)

## Preserving a low cardinality for the request counter

The request counter (`requests_total`) has a `url` label which,
although desirable, can become problematic in cases where your
application uses templated routes expecting a great number of
variations, as Prometheus explicitly recommends against metrics having
high cardinality dimensions:

https://prometheus.io/docs/practices/naming/#labels

If you have for instance a `/customer/:name` templated route and you
don't want to generate a time series for every possible customer name,
you could supply this mapping function to the middleware:

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zoftdev/go-gin-prometheus"
)

func main() {
	r := gin.New()

	p := ginprometheus.NewPrometheus("gin")

	p.ReqCntURLLabelMappingFn = func(c *gin.Context) string {
		url := c.Request.URL.Path
		for _, p := range c.Params {
			if p.Key == "name" {
				url = strings.Replace(url, p.Value, ":name", 1)
				break
			}
		}
		return url
	}

	p.Use(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello world!")
	})

	r.Run(":29090")
}
```

which would map `/customer/alice` and `/customer/bob` to their
template `/customer/:name`, and thus preserve a low cardinality for
our metrics.


## override http response status

if you use gin_request_* to report healthy, and want to override code field e.g. 
```
 gin_request_duration_seconds_bucket{code="2002",method="POST",url="/api/v1/updateStatus",le="5"} 1
```
use StatusOverrideFromContext
```

	p := ginprometheus.NewPrometheus("gin")
	p.StatusOverrideFromContext = append(p.StatusOverrideFromContext, "code")
	p.StatusOverrideFromContext = append(p.StatusOverrideFromContext, "mycode")

``` 

and set in gin's context
```
r.GET("/", func(c *gin.Context) {
		c.Set("code", "888")
		c.JSON(200, "Hello world!")
	})
```