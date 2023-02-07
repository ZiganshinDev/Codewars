/* Complete the solution so that it reverses all of the words within the string passed in.

Words are separated by exactly one space and there are no leading or trailing spaces.

Example(Input --> Output):

"The greatest victory is that which requires no battle" --> "battle no requires which that is victory greatest The" */

// My solution

package kata

import "strings"

func ReverseWords(str string) string {
	strarray := strings.Split(str, " ")
	var newstrarray []string
	for i := len(strarray) - 1; i >= 0; i-- {
		newstrarray = append(newstrarray, string(strarray[i]))
	}
	str = strings.Join(newstrarray, " ")
	return str // reverse those words
}
