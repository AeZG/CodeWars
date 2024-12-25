/*Write a function to convert a name into initials. This kata strictly takes two words with one space in between them.

The output should be two capital letters with a dot separating them.

It should look like this:

Sam Harris => S.H

patrick feeney => P.F*/

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
