
package main

import (
  "fmt"
  "os"
  "flag"
  "regexp"
)

var (
  help *bool = flag.Bool("h", false, "to show help")
  mode *bool = flag.Bool("m", false, "to run in std mode")
  runs *int = flag.Int("runs", 100000, "number of runs to do")
  re *string = flag.String("re", "a*a*a*a*a*aaaaa", "regexp to build")
  s *string = flag.String("s", "aaaaa", "string to match")
)

func main() {
  flag.Parse()
  if *help {
    flag.PrintDefaults()
    return
  }

  if !*mode {
    // use new regexp impl
    r := Parse(*re)

    for i := 0; i < len(r.prog); i++ {
      fmt.Fprintln(os.Stderr, i, r.prog[i].str())
    }

    result := false
    var alt []pair
    for i := 0; i < *runs; i++ {
      result, alt = r.run(*s)
    }

    fmt.Fprintln(os.Stderr, "new result", result, "alt", alt)
  } else {
    // use old regexp impl
    r := regexp.MustCompile(*re)
    var result []int
    for i := 0; i < *runs; i++ {
      result = r.ExecuteString(*s)
    }
    success := (len(result) != 0)
    if success {
      result = result[2:]
    }
    fmt.Fprintln(os.Stderr, "std result", success, "alt", result)
  }
}
