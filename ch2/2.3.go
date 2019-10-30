package ch2

/*
2.3 - 1
Illustrate the operation of merge sort on the array:
{3, 41, 52, 26, 38, 57, 9, 49}

Recursion Level 0 (Down):
L: {3, 41, 52, 26, 38, 57, 9, 49} R: {}

Recursion Level 1 (Down):
L: {3, 41, 52, 26} R: {38, 57, 9, 49}

Recursion Level 2 (Down):
L1: {3, 41} R1: {38, 57}
L2: {52, 26} R2: {9, 49}

Recursion Level 3 (Down):
L1: {3} R1:{41}
L2: {38} R2: {57}
L3: {52} R3: {26}
L4: {9} R4: {49}

Traverse back up the recursion tree - The "Conquer" part

Recursion Level 2 (Up):
L1: {3, 38} R1: {41, 57}
L2: {9, 52} R2: {26, 49}

Recursion Level 3 (Up):
L: {3, 9, 38, 52} R: {26, 41, 49, 57}

Recursion Level 0 (Up):
L: {3, 9, 26, 38, 41, 49, 52, 57} R:{}
*/

/*
2.3 - 2
Rewrite the merge sort so it does not use sentinels.

Instead of removing elements until we hit the sentinel card, we can
remove  elements until the array is empty.
*/

/*
2.3 - 3
Because it is a mathematical proof answer is omitted.
*/

/*
2.3 - 4
Write a recurrence relation for a recursive implementation for insertion sort.

T(n):
    T(1) = O(1) if n = 1
    T(n) = T(n - 1) + O(n) if n > 1
*/

/*
2.3 - 5
Write psuedocode for binary search. Argue that the worst case run time is O(log n)
*/

func binarySearch(A []int, left int, right int, v int) int {
	m := (left + right) / 2
	if A[m] == v {
		return m
	}

	if A[m] > v {
		return binarySearch(A, left, m-1, v)
	} else {
		return binarySearch(A, m+1, right, v)
	}
}

/*
The worst case for binary search is O(log n). This is because we effectively half the size of the array
every recursive iteration.
*/

/*
2.3 - 6
Can we use binary search to improve the performance of insertion sort?

No we cannot. We are not simply performing a linear scan on the previous elements, we are also moving them forward which is a O(n) operation.
Since we drop non-dominating terms, we drop the O(log n) since O(n) dominates it. This means that the time complexity has not improved. Also it
should be noted that using binary search in lieu of linear search would also make the code much more complicated to read.
*/

/*
2.3 - 7
Describe a O(n log n) algorithm that, given a set S of n intergers and another integer x, whether or not there exists a pair whose sum is exactly x.

There are two ways we can solve this. We can sort the array first and then use a while loop to check for pairs. Since sorting is a O(n log n) operation
and the check is a O(n) operation, this solution is a O(n log n) solution.

Code is written in C++ here

bool isMatch(vector<int>& S, int x) {
    sort(S, S.size());
    int i = 0, j = S.size() - 1;

    while(i <= j) {
        int target = x;
        if(S[i] + S[j] == x) return true;
        if(S[i] + S[j] > x) j--;
        else i++;
    }
    return false;
}

The other probably more theme appropriate solution is to perform binary search on the array for every element and check if it has the complement.
Since we are performing a binary search inside a linear scan our time complexity is O(n log n).

bool isMatch(vecotr<int>& S, int x) {
    sort(S, S.size());

    for(int i = 0; i < S.size(); i++) {
        int complement = x - S[i];
        if(binarySearch(S, i, S.size(), complement)) return true;
    }
    return false;
}
*/
