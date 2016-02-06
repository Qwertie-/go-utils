package main

import (
    "fmt"
    "os"
)

func main() {
  for {
    if len(os.Args) == 1 {
      fmt.Println("y")
    } else {
      for i := 1; i < len(os.Args); i++ {
        fmt.Print(os.Args[i] + " ") //Printing this way to get rid of the [] around arrays
      }
      fmt.Println()
    }
  }
}
