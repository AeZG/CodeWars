/*Build a function that returns an array of integers from n to 1 where n>0.

Example : n=5 --> [5,4,3,2,1]*/

package ReverseSequence

func ReverseSeq(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = n - i
	}
	return arr
}
