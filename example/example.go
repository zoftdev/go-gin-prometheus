package main

import (
	"fmt"

	ginprometheus "github.com/zoftdev/go-gin-prometheus"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	/*	// Optional custom metrics list
		customMetrics := []*ginprometheus.Metric{
			&ginprometheus.Metric{
				ID:	"1234",				// optional string
				Name:	"test_metric",			// required string
				Description:	"Counter test metric",	// required string
				Type:	"counter",			// required string
			},
			&ginprometheus.Metric{
				ID:	"1235",				// Identifier
				Name:	"test_metric_2",		// Metric Name
				Description:	"Summary test metric",	// Help Description
				Type:	"summary", // type associated with prometheus collector
			},
			// Type Options:
			//	counter, counter_vec, gauge, gauge_vec,
			//	histogram, histogram_vec, summary, summary_vec
		}
		p := ginprometheus.NewPrometheus("gin", customMetrics)
	*/

	p := ginprometheus.NewPrometheus("gin")
	p.StatusOverrideFromContext = append(p.StatusOverrideFromContext, "code")
	p.StatusOverrideFromContext = append(p.StatusOverrideFromContext, "mycode")
	p.Use(r)
	r.GET("/:hlex", func(c *gin.Context) {

		fmt.Println(c.FullPath())
		c.Set("code", "888")
		c.JSON(200, "Hello world!")
	})
	r.GET("/test2", func(c *gin.Context) {
		c.Set("mycode", "999")
		c.JSON(200, "Hello world!")
	})
	r.Run(":29090")
}
