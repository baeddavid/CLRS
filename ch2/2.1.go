package ch2

/*
2.1-1
Illustrate the operation of Insertion Sort on the array A = {31, 41, 59, 26, 41, 58}

A0 = {31, 41, 59, 26, 41, 58}
A1 = {31, 41, 59, 26, 41, 58}
A2 = {31, 41, 59, 26, 41, 58}
A3 = {26, 31, 41, 59, 41, 58}
A4 = {26, 31, 41, 41, 59, 58}
A5 = {26, 31, 41, 41, 58, 59}
*/

/*
2.1-2
Rewrite Insertion Sort to sort into nonincreasing instead of non-decreasing order.
*/

func iSort(a []int) {
	for i := 1; i < len(a); i++ {
		key := a[i]
		j := i - 1
		// The second part of the AND operator is what does the comparison for sorting into
		// increasing order. Simply flipping the direction of the operator yields us the
		// array sorted into nonincreasing order
		for j >= 0 && a[j] < key {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = key
	}
}

/*
2.1-3
Consider the following searching problem
Input: A sequence of numbers such that A = {a1, a2, ..., an} and a value of v
Ouput: An index i such that v = A[i] or the spcial value NIL if v ∉ A


Write code for Linear Search which scans through the sequence. Using loop invariant
prove that the algorithm is correct.
*/

func linearSearch(A []int, v int) (int, bool) {
	for idx, val := range A {
		if val == v {
			return idx, true
		}
	}
	// Our NIL return
	return -1, false
}

/*
Initialization:
We know that all values prior to the first loop are not equal to the key. Before we start looping we
can return NIL which is a valid return

Maintenance:
Our idx variable will hold NIL until either we terminate out of the loop or we find a value equal to v

Termination:
When we exit the loop we have either NIL or idx where A[idx] is equal to v
*/

/*
2.1-4
Consider the problem of adding two n-bit binary integers, stored in two n-element arrays A and B. The sum
of the two integers should be stored in binary form in an (n + 1) element array C. State the problem formally
and write the code adding the two integers.
*/

func addTwo(A []int, B []int) []int {
	C := make([]int, len(A)+1)
	carry := 0
	for i := 0; i < len(C); i++ {
		C[i] = (A[i] + B[i] + carry) % 2
		carry = (A[i] + B[i] + carry) / 2
	}
	C[len(C)-1] = carry
	return C
}
