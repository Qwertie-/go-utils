package main

import (
  "time"
  "os"
  "reflect"
  "fmt"
)

func main() {
  if reflect.TypeOf(os.Args[1]) == "int" or reflect.TypeOf(os.Args[1]) == "float" {
    sleep(os.Args[1])
  } else {
    fmt.Println("sleep: invalid time interval '" + os.Args[1] + "'")
  }
}
