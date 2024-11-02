package main

import (
  "fmt"
  "github.com/cahyasetya/go-channel-pattern/generator"
)

func main() {
  result := generator.Sum(3, 4)
  fmt.Printf("result: %d", <-result)
}
