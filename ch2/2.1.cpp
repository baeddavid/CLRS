using namespace std;
#include<vector>
/*
2.1 - 2
Rewrite Insertion-Sort to sort into nonincreasing instead of non-decreasing order.
*/
void decreaseInsertionSort(vector<int>& V) {
    for(int i = 1; i < V.size(); i++) {
        int key = V[i];
        int j = i - 1;
        // We simply flip the comparison in the second half of the while check
        while(j > -1 && V[j] < key) {
            V[j + 1] = V[j];
            j--;
        }
        V[j + 1] = key;
    }
}

/*
2.2 - 3
Consider the searching problem:
    Input: A sequence of n numbers A = {a1, a2, ..., an} and a value v.
    Output: An index i such that v = A[i] or the special value NIL if v does not appear in A.
Write pseudocode for linear search, which scans through the sequence, looking for v. Using a loop
invariant, prove that your algorithms is correct. 
*/

int linearSearch(vector<int>& V, int v) {
    for(int i = 0; i < V.size(); i++)
        if(V[i] == v)
            return i;
    return NULL;
}

/*
Initialization: 
    We know that all values prior to index i are not equal to our key. This holds prior to initialization 
    as before we start looping, we know that the element at index i (which has not been initialized yet) is not 
    equal to the key.

Maintenance: 
    As we loop through the array, we know that all elements prior to the element at index i is not the key. 

Termination:
    We terminate in two cases. Either we find a value at index i that is equal to v and we can immediately return 
    or we reach the end of the sequence and return NIL.
*/

/*
2.1 - 4
Consider the problem of adding two n-bit integers, stored in two n-element arrays A and B. The sum of the two intergers
should be stored in binary form in an (n + 1)-element array C. State the problem formally and write pseudocode for adding 
two integers.

Given an array A that represents a value in an n bit array and an array B that also represents an n bit array, return an
array C that is the sum of A and B in binary form in an array with n + 1 elements.

addTwo(A, B) {
    C = new int[A.size + 1]
    carry = 0;
    for(i = 0 to C.sie) {
        C[i] = (A[i] + B[i] + carry) % 2
        carry = (A[i] + B[i] + carry) / 2
    }
    C[C.size - 1] = carry
    return C
}
*/