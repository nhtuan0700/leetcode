def mergeSort(arr, l, r):
  if l < r:
    m = l + (r - l) // 2
    
    mergeSort(arr, l, m)
    mergeSort(arr, m + 1, r)
    merge2(arr, l, r)

def merge(arr, l, m, r):
  # size of list 1
  n1 = m - l + 1
  # size of list 2
  n2 = r - m

  # copy data to list 1
  L1 = arr[l:l+n1]
  # copy data to list 2
  L2 = arr[m+1:m+1+n2]
    
  # sorted temp arry
  tmp = [0] * (n1 + n2)
  i, j = 0, 0
  k = 0
  while i < n1 and j < n2:
    if L1[i] <= L2[j]:
      tmp[k] = L1[i]
      i += 1
    else:
      tmp[k] = L2[j]
      j += 1
    k += 1
  
  # Copy remaingng elements of L1 if any
  while i < n1:
    tmp[k] = L1[i]
    i += 1
    k += 1
  # Copy remaingng elements of L2 if any
  while j < n2:
    tmp[k] = L2[j]
    j += 1
    k += 1
  
  for i in range(len(tmp)):
    arr[l+i] = tmp[i]
    
    
def merge2(arr, l, r):
  m = l + (r - l) // 2
  n = r-l+1
  tmp = [0] * n
  i, j = l, m+1
  for k in range(0, n):
    if j > r or (i <= m and arr[i] <= arr[j]):
      tmp[k] = arr[i]
      i += 1
    else:
      tmp[k] = arr[j]
      j += 1
  
  for i in range(0, n):
    arr[i+l] = tmp[i]

arr = [10,4,5,1,2,7,4]
mergeSort(arr, 0, len(arr) -1)
print(arr)
