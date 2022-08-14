# # This function is to implement the binary search
arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

def binary_search(arr, value):
 low = 0
 high = len(arr) - 1
 while low <= high:
  midpoint = (low + high)
  guess = arr[midpoint]
  if guess == value:
   return f"Mid point {midpoint} with value {value} found!"
  elif guess < value:
   low = midpoint + 1
  else:
   high = midpoint - 1
 print("Number does not exist!\n")
 return None

print(binary_search(arr, 3))
