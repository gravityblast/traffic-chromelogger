/*
Package chromemiddleware implements a Chrome Logger middleware for Traffic.

This is a Middleware for traffic (https://github.com/pilu/traffic).
It allows to send logs to the Chrome console if you have the  Chrome Logger extension (http://craig.is/writing/chrome-logger) installed.

Usage:

  package main

  import (
    "log"
    "fmt"
    "net/http"
    "github.com/pilu/traffic"
    "github.com/pilu/traffic-chromelogger"
  )

  func rootHandler(w traffic.ResponseWriter, r *http.Request) {
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

    http.Handle("/", router)
    err := http.ListenAndServe(":7000", nil)

    if err != nil {
      log.Fatal(err)
    }
  }
*/
package chromelogger
