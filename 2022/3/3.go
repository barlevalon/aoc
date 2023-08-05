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
    chars := []rune(scanner.Text())
    firstHalf := make(map[rune]int)
    for _,char := range chars[:(len(chars)/2)] {
      firstHalf[char] = 1
    }
    secondHalf := chars[len(chars)/2:]
    for _, char := range secondHalf {
      _, ok := firstHalf[char]
      if ok {
        sum += priorities[char]
        break
      }
    }
  }
  fmt.Println(sum)
}
