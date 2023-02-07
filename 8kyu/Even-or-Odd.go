/* Create a function that takes an integer as an argument and returns "Even" for even numbers or "Odd" for odd numbers. */

// My solution

package kata

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}
