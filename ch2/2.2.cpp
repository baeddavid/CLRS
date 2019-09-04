#include<vector>
using namespace std;

/*
2.2 - 1
n^3 / 1000 - 100n^2 - 100n + 3

When we calculate time complexities we drop constants and non-dominating terms.
This means we can drop everything but the n^3 since it is the dominating term.

Answer: O(n^3)
*/

/*
2.2 - 2
Write pseudocode for selection sort, loop invariant for it, and the time best,
average, and worst case time complexities.

Selection sort pseudocode:
for(i = 0 to A.l - 1)
    min = i
    for(j = i + 1 to A.l)
        if(A[j] < A[min])
            min = j
    swap(A[min], A[i])

Selection sort implementation:
*/
void selectionSort(vector<int>& A) {
    for(int i = 0; i < A.size(); i++) {
        int min = i;
        for(int j = i + 1; j < A.size(); j++) {
            if(A[j] < A[min])
                min = j;
        }
        int temp = A[i];
        A[i] = A[min];
        A[min] = temp;
    }
}

/*
Worst Case: O(n^2)
Best Case: O(n^2)

Loop invariant:
    Initialization -> Prior to starting the loop, i - 1 contains all elements originally
    in A but in sorted order. Since i = 0, i - 1 indicates that there are no sorted elements
    prior to the start of the sort.

    Maintenance -> At every iteration of the loop, i - 1 contains the smallest to i - 1 
    elements in A.

    Termination -> The loop terminates when i = A.l. When i = A.l we have terminated out
    of the loop and we can infer that A is now in sorted order.
*/

/*
2.2 - 3
What is the Average case and Worst case for linear search?

Worst case is obviously O(n) where n is the length of the array.
Average case is a little more complicated. We can make the assumption that for every iteration
of the loop, we have a 1 / n chance to find the key. We can rephrase this as every key has a 
1 / n chance of being the key
*/

/*
2.2 - 4
How can we modify any algorithm to have a good Best case time complexity?

Randomize an output and then check if it is correct. It is O(n) time for best case.
*/