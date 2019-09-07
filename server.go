package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
)

func getIp() string {
	ifaces, _ := net.Interfaces()
	// handle err
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			// process IP address
			if !ip.IsLoopback() {
				return ip.String()
			}
		}
	}
	return "localhost"
}

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	addr := getIp() + ":8080"
	fmt.Println(addr)
	flag.Parse()
	log.SetFlags(0)
	directory := flag.String("d", ".", "")
	http.Handle("/", http.FileServer(http.Dir(*directory)))
	log.Fatal(http.ListenAndServe(addr, nil))
}
