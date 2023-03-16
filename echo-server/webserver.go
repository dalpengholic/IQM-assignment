package main

import (
	"flag"
	"fmt"
	"net/http"
        "os"
)

// Define struct to hold data of `flag`
type EchoConfig struct {
	port_number string
	wel_message string
	hostname string
}

// Initialization of variables
func InitConfig(param EchoConfig) EchoConfig {
	flag.StringVar(&param.port_number, "p", "8000", "default port number")
	flag.StringVar(&param.wel_message, "c", "hello", "default message")
	param.hostname, _ = os.Hostname()
	flag.Parse()

	return param
}

func main() {
	var holder EchoConfig
	holder = InitConfig(holder)
	holder.port_number = ":" + holder.port_number

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, holder.wel_message, "from", holder.hostname)
	})

	http.ListenAndServe(holder.port_number, nil)
}
