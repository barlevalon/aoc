package main

import (
  "os"
  "fmt"
  "bufio"
  "strconv"
  "sort"
)

func checkErr(err error) {
  if err != nil {
    panic(err)
  }
}

func main() {

  file, err := os.Open("input")
  checkErr(err)
  defer file.Close()

  scanner := bufio.NewScanner(file)
  // var calories []int
  curr := 0
  var calories []int
  for i := 0; scanner.Scan(); {
    line := scanner.Text()
    if line != "" {
      lineInt, err := strconv.Atoi(line)
      checkErr(err)
      curr += lineInt
    } else {
      calories = append(calories, curr)
      curr = 0
      i++
    }
  }
  sort.Ints(calories)
  var sum int
  for i:=1; i <= 3; i++ {
    sum += calories[len(calories)-i]
  }
  fmt.Println(sum)
}
