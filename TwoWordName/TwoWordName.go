package TwoWordName

import (
  "fmt"
  "strings"
)

func AbbrevName(name string) string {
  splitName := strings.Split(name, " ")
  initial := strings.ToUpper(fmt.Sprintf("%s.%s", splitName[0][:1], splitName[1][:1]))
  return initial
}
