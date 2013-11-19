package main

import (
  "fmt"
  "github.com/pilu/traffic"
  "github.com/pilu/traffic-chromelogger"
)

func rootHandler(w traffic.ResponseWriter, r *traffic.Request) {
  logger := w.GetVar("chromelogger").(*chromelogger.Logger)

  logger.Log("Hello")
  logger.Log(map[string]string{
    "foo": "bar",
  })

  fmt.Fprint(w, "Hello, check your Chrome console after activating the Chrome Logger extension.\n")
}

func main() {
  router := traffic.New()
  router.AddMiddleware(chromelogger.New())
  router.Get("/", rootHandler)
  router.Run()
}
