package main

import (
	"encoding/json"
	"module06/app/utils"
	"net/http"
)

type Response struct {
	Msgs []string
}

func handler(w http.ResponseWriter, r *http.Request) {
	resp := Response{Msgs: make([]string, 10)}

	for i := 0; i < 1000; i++ {
		msg := utils.Pad("Hello, ", i, " World!")
		resp.Msgs = append(resp.Msgs, msg)
	}

	jsonRes, _ := json.Marshal(resp)
	w.Write(jsonRes)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9090", nil)
}
