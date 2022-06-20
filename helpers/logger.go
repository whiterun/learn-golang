package helpers

import (
    "log"
    "os"
    "sync"
)

type initial struct {
    filename string
    *log.Logger
}

var logger *initial
var once sync.Once

func LogInstance() *initial {
    once.Do(func() {
        logger = createLogger("./tmp/app.log")
    })
    return logger
}

func createLogger(fname string) *initial {
    file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

    return &initial{
        filename: fname,
        Logger: log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
    }
}