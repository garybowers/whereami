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

	http.HandleFunc("/", root)
	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func root(w http.ResponseWriter, req *http.Request) {
	clusterName, _ := getCurrentClusterName()
	clusterLocation, _ := getCurrentClusterLocation()

	log.Println(clusterName + " -- " + clusterLocation)

	fmt.Fprintf(w, clusterName)
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, clusterLocation)
}

func getCurrentClusterName() (clusterName string, err error) {
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

func getCurrentClusterLocation() (clusterLocation string, err error) {
	url := "http://metadata/computeMetadata/v1/instance/attributes/cluster-location"
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
