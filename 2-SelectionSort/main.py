from time import time


def findSmallest(arr):
 smallest = arr[0] # store the smallest value - which starts as the first value
 smallest_index = 0 # store the smallest index of the smallest value
 for i in range(1, len(arr)):
  if arr[i] < smallest:
   smallest = arr[i]
   smallest_index = i
  
 return smallest_index

def selectionSort(arr):
 newArr = []
 for i in range(len(arr)):
  smallest = findSmallest(arr)
  newArr.append(arr.pop(smallest))
 
 return newArr

arr = [6, 2, 5, 3, 8, 1]
print("Old array: ", arr)
print("New array: ", selectionSort(arr))