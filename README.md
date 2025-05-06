Greetings to everyone,

In this article, I will try to explain the most famous data search algorithms in Go language.

**_Linear Search_**

All array elements are checked sequentially, starting from the first element. The process ends when the searched number is found; otherwise, the entire array is scanned.

The complexity of linear search is O(n).

Code in Go-->



![Linear Search](/images/LinearSearch.JPG)



**_Binary Search_**

Number sequence(array) must be sequential. The target number is compared with the middle element of the array, and if it is equal to the middle element, the operation is successful and the fastest result is obtained.
If it is not equal to the middle element, it is checked.
    

- If it is less than the middle element, a new search is started in the left half of the array.
-     If it is larger than the middle element, the search starts in the right half of the array.



Best Case Time Complexity of Binary Search Algorithm: O(1)

The average case complexity is O(logN)
N-->length

Code in Go


![Binary Search](/images/BinarySearch.JPG)



**_Sentinel Linear Search_**

The complexity of Sentinel Linear Search O(n).

The number of comparisons are less in _sentinel linear_ search than _linear search_


![SentinelSearch](/images/SentinelSearch.JPG)



**_Jump Search_**

Number sequence(array) must be sequential
The complexity of Sentinel Jump Search is O(√N).(Worst Case)



![JumpSearch](/images/JumpSearch.JPG)



-----------------------------------------------
