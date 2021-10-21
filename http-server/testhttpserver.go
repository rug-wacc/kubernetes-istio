package main

import (
	"fmt"
	"net"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func ip(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "----------------------------Response Start-----------------\n")
	list, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for i, iface := range list {
		fmt.Fprintf(w, "Response %d name=%s %v\n", i, iface.Name, iface)
		addrs, err := iface.Addrs()
		if err != nil {
			panic(err)
		}
		for j, addr := range addrs {
			fmt.Fprintf(w, "Response %d-%d %v\n", i, j, addr)
		}
	}
	fmt.Fprintf(w, "----------------------------Response End-----------------\n")
	fmt.Printf("Test")
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/ip", ip)
	http.ListenAndServe(":9099", nil)
}
