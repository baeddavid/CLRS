package ch2

/*
2.2-1
Express the function n^3 / 1000 - 100n^2 - 100n + 3 in tems of O notation

We drop all constants leaving and nondominating terms leaving us with O(n^3)
*/

/*
2.2-2
Write code for Selection Sort. What loop invariant those this algorithm maintain? Why does it need to only
run for n - 1 elements, rather than for all n element? Give the best-case and worst-case running times of
Selection Sort in O notation
*/

func selectionSort(A []int) {
	for i := 0; i < len(A)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(A); j++ {
			if A[j] < A[minIdx] {
				minIdx = j
			}
		}
		A[i], A[minIdx] = A[minIdx], A[i]
	}
}

/*
Initialization: Prior to starting the loop, i - 1 contains all elements originally
in A but in sorted order. Since i = 0, i - 1 indicates that there are no sorted elements
prior to the start of the sort.

Maintenance: At every iteration of the loop, i - 1 contains the smallest to i - 1
elements in A.

Termination: The loop terminates when i = A.l. When i = A.l we have terminated out
of the loop and we can infer that A is now in sorted order.

Worst Case: O(n^2)
Best Case: O(n^2)
*/

/*
2.2-3
Consider Linear Search Again. How many elements of the input sequence need to be checked on average assuming that the element being
searched for is equally likely to be any element in the array? How about in worst case? What are the average time and worst time in O
notation?

Obviously worst case is when v is not present in the set A. We will have to check every element in the array only to return NIL meaning
that Linear Search has a worst time complexity of O(n).

The average case is a little bit more complicated. When searing for an element in a set the value v may be located at postion 1, 2, ... or n.
This means that when we reach the kth element in the array we have made k comparisons.

To average this out we know that 1 + 2 + ... + n is equal to (n + 1)n / 2. We then divide it by n to get the average. This yields
n + 1 / 2. We then drop constants and nondominating terms getting O(n) for our average case.

Worst Case: O(n)
Average Case: O(n)
*/

/*
2.2-4
How can we modify any algorithm to have a good best-case running time?

We can randomly generate an output corresponding to the return type of the algorithm and check if it is valid. If valid then return it.
This is a O(n) possible solution.
*/
