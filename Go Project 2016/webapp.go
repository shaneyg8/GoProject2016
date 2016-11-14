
// Adapted from: https://go-macaron.com

package main

import "gopkg.in/macaron.v1"

// Function taken from: github.com/golang/example


func main() {
  m := macaron.Classic()
  m.Use(macaron.Renderer())
  
  m.Get("/macaron", func() string {
    return "Hello from Macaron!"
  })
  
  m.Run(8080)
}
