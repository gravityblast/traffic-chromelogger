package main

import (
  "github.com/pilu/traffic"
  "github.com/pilu/traffic-chromelogger"
)

func rootHandler(w traffic.ResponseWriter, r *traffic.Request) {
  logger := w.GetVar("chromelogger").(*chromelogger.Logger)

  logger.Log("Hello")
  logger.Log(map[string]string{
    "foo": "bar",
  })

  w.WriteText("Hello, check your Chrome console after activating the Chrome Logger extension.\n")
}

func main() {
  router := traffic.New()
  router.Use(chromelogger.New())
  router.Get("/", rootHandler)
  router.Run()
}
