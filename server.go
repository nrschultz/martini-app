package main

import "github.com/go-martini/martini"
import "fmt"

func main() {
  m := martini.Classic()
  m.Get("/", func() string {
    ages := map[string]int{"nick": 26, "ruth": 28}
    fmt.Println("map:", ages)
    fmt.Println("map:", ages["mark"])
    return "Hello world!"
  })
  m.Run()
}
