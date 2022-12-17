package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	fmt.Println("Where Am I? v0.0.1")
	fmt.Println("==================")

	clusterName, _ := getCurrentCluster()
	log.Println(clusterName)

}

func getCurrentCluster() (clusterName string, err error) {
	url := "http://metadata/computeMetadata/v1/instance/attributes/cluster-name"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Metadata-Flavor", "Google")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	b, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatal(err)
	}

	return string(b), nil
}
