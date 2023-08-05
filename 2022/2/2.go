package main

import(
  "os"
  "bufio"
  "strings"
  "fmt"
)

func checkErr(err error) {
  if err != nil {
    panic(err)
  }
}


var scores = [3][3]int {
  {3, 4, 8},
  {1, 5, 9},
  {2, 6, 7},
}
func score(split []string) (score int) {
  opponent := split[0]
  outcome := split[1]
  var x int
  switch opponent {
    case "A": x = 0
    case "B": x = 1
    case "C": x = 2
  }
  var y int
  switch outcome {
    case "X": y = 0
    case "Y": y = 1
    case "Z": y = 2
  }
  return scores[x][y]
}

func main() {

  file, err := os.Open("input")
  checkErr(err)
  defer file.Close()

  scanner := bufio.NewScanner(file)
  var total int
  for scanner.Scan() {
    line := scanner.Text()
    split := strings.Split(line, " ")
    total += score(split)
  }
  fmt.Println(total)
}
