package main

import (
  "os"
  "fmt"
  "io/ioutil"

  "github.com/charmbracelet/glamour"
)

func main() {
  if len(os.Args) < 2 {
    fmt.Print("usage: mdrender filepath\n")
    os.Exit(1)
  }

  filepath := os.Args[1]

  b, err := ioutil.ReadFile(filepath)

  if err != nil {
    fmt.Print(err)
    os.Exit(1)
  }

  tr, err := glamour.NewTermRenderer(glamour.WithEnvironmentConfig())

  if err != nil {
    fmt.Print(err)
    os.Exit(1)
  }

  str := string(b)

  parsed, err := tr.Render(str)

  if err != nil {
    fmt.Print(err)
    os.Exit(1)
  }

  parsed = fmt.Sprintf("%s\n\n", parsed)

  fmt.Fprint(os.Stdout, parsed)
}

