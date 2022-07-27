package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

var (
	addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
)

var (
	testCountVector = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "test_count_vector",
			Help: "",
		},
		[]string{"label1", "label2"},
	)
)

func init() {
	prometheus.MustRegister(testCountVector)

}

func runServer() {
	flag.Parse()

	go func() {
		fmt.Println("add metric series with label1=v1, label2=v2")
		testCountVector.WithLabelValues([]string{"v1", "v2"}...).Set(1)
		time.Sleep(10 * time.Second)
		fmt.Println("delete metric series with label1=v1, label2=v2")
		testCountVector.DeleteLabelValues([]string{"v1", "v2"}...)
	}()

	// Expose the registered metrics via HTTP.
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func runSeries() {

	client, err = api.NewClient(api.Config{
		Address: "http://47.112.190.225:9090/",
	})
	
	if err != nil {
		return cli, err
	}

	cli = v1.NewAPI(client)

	matches := []string{"etcd_cluster_info"} 

	lbs, _, err := cli.Series(ctx, matches, Add(time.Duration(-60)*time.Second), time.Now())	
	if err != nil {
		fmt.Println("err", err)
		return
	}

	fmt.Println(lbs)
}

func main() {

	runSeries()
}


