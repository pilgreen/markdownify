package main

import (
  "bytes"
  "os"
  "io"
  "io/ioutil"
  "log"

  "github.com/russross/blackfriday"
)

func main() {
  fi, _ := os.Stdin.Stat()
  if fi.Mode() & os.ModeNamedPipe != 0 {
    b, err := ioutil.ReadAll(os.Stdin)
    if err != nil {
      log.Fatalf("%s .. error reading Stdin", os.Args[0])
    }

    reader := bytes.NewBuffer(blackfriday.MarkdownCommon(b))
    io.Copy(os.Stdout, reader)
  }
}
