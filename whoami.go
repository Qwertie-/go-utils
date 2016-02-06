package main

import (
  "os/user"
  "fmt"
)

func main() {
  u, err := user.Current()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.Username)
}
