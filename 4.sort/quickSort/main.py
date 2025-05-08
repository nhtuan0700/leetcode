import random

def quickSort(arr, l, r):
    if l > r: return

    pI = partition3(arr, l, r)
    quickSort(arr, l, pI - 1)
    quickSort(arr, pI + 1, r)

# select last element for pivot
def partition(arr, l, r):
    pivot = arr[r]
    i = l
    for j in range(l, r):
        if arr[j] < pivot:
            arr[i], arr[j] = arr[j], arr[i]
            i += 1

    arr[i], arr[r] = arr[r], arr[i]  
    return i

# select random
def partition2(arr, l, r):
    p = random.randint(l, r)
    arr[p], arr[r] = arr[r], arr[p]
    i = l
    for j in range(l, r):
        if arr[j] < arr[r]:
            arr[j], arr[i] = arr[i], arr[j]
            i += 1
    
    arr[i], arr[r] = arr[r], arr[i]
    return i

def partition3(arr, l, r):
    p = random.randint(l, r)
    pivot = arr[p]
    i, j = l, r
    while i < j:
        while arr[i] < pivot:
            i += 1
        while arr[j] > pivot:
            j -= 1
        
        if i < j:
            arr[i], arr[j] = arr[j], arr[i]
            i += 1
            j -= 1

	# j <= i, j có thể là phần tử cuối cùng <= pivot hoặc là cái đầu tiên mà nó > pivot
	# [5 4 6 3 1 2], pivot = 6
    if arr[j] > pivot:
        j -= 1
    return j


arr = [10,4,5,1,2,7,4]
quickSort(arr, 0, len(arr) -1)
print(arr)
