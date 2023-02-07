/* Complete the solution so that it reverses the string passed into it.

'world'  =>  'dlrow'
'word'   =>  'drow' */

// My solution

package kata

import "strings"

func Solution(word string) string {
	var rs []string
	for i := len(word) - 1; i >= 0; i-- {
		rs = append(rs, string(word[i]))
	}
	justString := strings.Join(rs, "")
	return justString
}
