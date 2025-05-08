from typing import List

def bubbleSort(arr: List[int]):
    n = len(arr)
    for i in range(0, n-1):
        for j in range(0, n-i-1):
            if arr[j] > arr[j+1]:
                arr[i], arr[j] = arr[j], arr[i]
                

arr = [1,5,4,8,6,7]
bubbleSort(arr)

print(arr)

# [0,1,2,3]
# i = 2
# j: 0 -> 4-1-2
