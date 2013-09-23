package main

import (
  "log"
  "fmt"
  "net/http"
  "github.com/pilu/traffic"
  "github.com/pilu/chromelogger"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
  chromelogger.Logger.Log("Hello")

  chromelogger.Logger.Log(map[string]string{
    "foo": "bar",
  })

  fmt.Fprint(w, "Hello World\n")
}

func main() {
  router := traffic.New()
  router.AddMiddleware(chromelogger.New())
  router.Get("/", rootHandler)

  http.Handle("/", router)
  err := http.ListenAndServe(":7000", nil)

  if err != nil {
    log.Fatal(err)
  }
}
