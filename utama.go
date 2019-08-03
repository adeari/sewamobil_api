package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
)

var port string = "8083"
func main() {
  http.HandleFunc("/", homePage)
  fmt.Println(port)
  log.Fatal(http.ListenAndServe(":"+port, nil))
}

func homePage(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Access-Control-Allow-Origin", "*")
  
    urlny := r.URL.Path
    if len(urlny) > 1 {
      urlny = urlny[1:len(urlny)]
    }    
    
    switch urlny {
      case "/":
        fmt.Fprintf(w, "")
      case "kota":
        json.NewEncoder(w).Encode(ambilKota())
      default:
        fmt.Println("ksong")
    }
}

type Kotapilihan struct {
  Id int `json:"id"`
  Kota string `json:"kota"`
}
func ambilKota()  []Kotapilihan {
  var kotas []Kotapilihan
  kotas = append(kotas, Kotapilihan{Id:0,Kota:"aaa"})
  kotas = append(kotas, Kotapilihan{Id:1,Kota:"bbb"})
  
  return kotas
}


