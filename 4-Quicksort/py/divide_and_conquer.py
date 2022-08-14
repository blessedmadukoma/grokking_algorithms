# Explaining divide and conquer concept
from time import process_time_ns

arr = [1,5,2,3,4, 8]

# def sum(arr):
#  start = process_time_ns()

#  sum = 0
#  for x in arr:
#   sum += x
 
#  stop = process_time_ns()
 
#  print(f"Took {stop - start}ns")
#  return sum

# print(sum(arr))

def recursive_sum(arr):
 start = process_time_ns()
 if arr == []:
  return 0
  
 stop = process_time_ns()
 print(f"Took {stop - start}ns")
 return arr[0] + recursive_sum(arr[1:])
 
# def rec_count_sum(arr):
#  if arr == []:
#   return 0
#  return 1 + rec_count_sum(arr[1:])
  
# def normal_max(arr):
#  max = 0
#  for x in arr:
#   if arr[x] > max:
#    max = arr[x]
#  return max
 
# def recursive_max(arr):
#  if len(arr) == 2:
#   return arr[0] if arr[0] > arr[1] else arr[1]
#  sub_max = recursive_max(arr[1:])
#  return arr[0] if arr[0] > sub_max else sub_max

print(f"Sum using recusion: {recursive_sum(arr)}")
# print(f"Count no of items using recursion: {rec_count_sum(arr)}")
# print(f"Max no using recursion: {recursive_max(arr)}")