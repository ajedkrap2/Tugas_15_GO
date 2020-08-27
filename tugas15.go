package main

import "fmt"
import "net/http"

var host = ":4848"

func loop(web http.ResponseWriter, req *http.Request){
  for i := 1 ; i <= 100 ; i++ {
    fmt.Fprintln(web, i)
  }
}

func main() {
    http.HandleFunc("/", func(web http.ResponseWriter, req *http.Request){
      fmt.Fprintln(web, "add /loop to web address to calculate")
    })

    http.HandleFunc("/loop", loop)

    fmt.Println("Memulai Web Server pada localhost", host)
    http.ListenAndServe(host, nil)
}
