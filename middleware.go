package chromelogger

import (
  "github.com/pilu/traffic"
)

type ChromeLoggerMiddleware struct {}

func (middleware ChromeLoggerMiddleware) ServeHTTP(w traffic.ResponseWriter, r *traffic.Request, next traffic.NextMiddlewareFunc) {
  logger := newLogger()

  rw := &responseWriter{
    ResponseWriter: w,
    logger:         logger,
    flushed:        false,
  }

  rw.SetVar("chromelogger", logger)

  if nextMiddleware := next(); nextMiddleware != nil {
    nextMiddleware.ServeHTTP(rw, r, next)
  }

  if !rw.flushed {
    rw.flush()
  }
}

func New() *ChromeLoggerMiddleware {
  middleware := &ChromeLoggerMiddleware{}

  return middleware
}
