package main

import (
	"fmt"
	"os"
  "bufio"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {

  priorities := make(map[rune]int)
  for i,char := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" {
    priorities[char] = i+1
  }

  file, err := os.Open("input")
  check(err)

  defer file.Close()

  scanner := bufio.NewScanner(file)

  sum := 0
  for scanner.Scan() {
    fmt.Println(scanner.Text())
    chars1 := []rune(scanner.Text())
    scanner.Scan()
    fmt.Println(scanner.Text())
    chars2 := []rune(scanner.Text())
    scanner.Scan()
    fmt.Println(scanner.Text())
    chars3 := []rune(scanner.Text())
    found := false
    for _, char := range chars1 {
      for _, char2 := range chars2 {
        for _, char3 := range chars3 {
          if char == char2 && char2 == char3 && !found {
            fmt.Println(string(char))
            sum += priorities[char]
            found = true
          }
        }
      }
    }
  }
  fmt.Println(sum)
}
