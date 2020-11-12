"""
Problem:
  Implement a priority queue backed by an array.
Constraints:
  Do we expect the methods to be insert, extract_min, and decrease_key?Yes
  Can we assume there aren't any duplicate keys?Yes
  Do we need to validate inputs?No
  Can we assume this fits memory?Yes
TestCases:
  insert general case -> inserted node
  extract_min from an empty list -> None
  extract_min general case -> min node
  decrease_key an invalid key -> None
  decrease_key general case -> updated node
  
Algorithm:
  insert : insert to the internal array.
  extract_min : 
    Loop through each item in the internal array
      Update the min value as needed
    Remove the min element from the array and return it
  decrease_key :
    Loop through each item in the internal array to find the matching input
      Update the matching element's key
"""

import sys

class PriorityQueueNode(object):
  def __init__(self, obj, key):
    self.obj = obj
    self.key = key
  
  def __repr__(self):
    return str(self.obj) + ": " + str(self.key)

class PriorityQueue(object):
  def __init__(self):
    self.array = []
  def __len__(self):
    return len(self.array)
  
  def insert(self, node):
    self.array.append(node)
    return self.array[-1]
  
  def extract_min(self):
    if not self.array:
      return None
    min = sys.maxsize
    for index, node in enumerate(self.array):
      if node.key < min:
        min = node.key
        min_index = index
    return self.array.pop(min_index)
  
  def decrease_key(self, obj, new_key):
    for node in self.array:
      if node.obj is obj:
        node.key = new_key
        return node
    return None

import unittest
class TestPriorityQueue(unittest.TestCase):
  def test_priority_queue(self):
    priority_queue = PriorityQueue()
    self.assertEqual(priority_queue.extract_min(), None)
    priority_queue.insert(PriorityQueueNode('a', 20))
    priority_queue.insert(PriorityQueueNode('b', 5))
    priority_queue.insert(PriorityQueueNode('c', 15))
    priority_queue.insert(PriorityQueueNode('d', 22))
    priority_queue.insert(PriorityQueueNode('e', 40))
    priority_queue.insert(PriorityQueueNode('f', 3))
    priority_queue.decrease_key('f', 2)
    priority_queue.decrease_key('a', 19)
    mins = []
    while priority_queue.array:
      mins.append(priority_queue.extract_min().key)
    self.assertEqual(mins, [2, 5, 15, 19, 22, 40])
    print('Success: test_min_heap')

def main():
  test = TestPriorityQueue()
  test.test_priority_queue()
if __name__ == '__main__':
  main()

