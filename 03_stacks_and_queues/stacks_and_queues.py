from typing import Optional, Union, Any, List
from collections import deque


class MyStack:
    def __init__(self) -> None:
        self._body: List[int] = []

    def __len__(self) -> int:
        return len(self._body)

    def push(self, item: int) -> None:
        self._body.append(item)

    def peek(self) -> int:
        if len(self._body) == 0:
            raise IndexError
        return self._body[-1]

    def pop(self) -> int:
        if len(self._body) == 0:
            raise IndexError
        return self._body.pop()


class MyQueue:
    def __init__(self) -> None:
        self._body: deque = deque()

    def __len__(self) -> int:
        return len(self._body)

    def add(self, item: Any) -> None:
        self._body.append(item)

    def remove(self) -> Any:
        if len(self._body) == 0:
            raise IndexError
        return self._body.popleft()

    def peek(self) -> Any:
        if len(self._body) == 0:
            raise IndexError
        return self._body[0]


## Stack via linked list
class Node:
    def __init__(self, data: int, is_end=False) -> None:
        self.data = data
        self.is_end = is_end
        self.next_node: Node


class LinkedListStack:
    def __init__(self) -> None:
        self._length = 0
        self._top = Node(0, is_end=True)

    def __len__(self) -> int:
        return self._length

    def push(self, item: int) -> None:
        new_node = Node(item)
        new_node.next_node = self._top
        self._top = new_node
        self._length += 1

    def pop(self) -> int:
        if self._top.is_end:
            raise IndexError
        else:
            ret = self._top.data
            self._top = self._top.next_node
            self._length -= 1
            return ret

    def peek(self) -> int:
        if self._top.is_end:
            raise IndexError
        else:
            return self._top.data


"""
3.1 Three in One:

Describe how you could use a single array to implement three stacks. 
"""


class FixedMultiStack:
    pass


"""
3.2 Stack Min:

How would you design a stack which, in addition to push and pop, has a function min
which returns the minimum element? Push, pop and min should all operate in 0(1) time. 
"""


class StackMin(MyStack):
    def __init__(self) -> None:
        super().__init__()
        self._min_tracker: List[int] = []

    def push(self, new_node: int) -> None:
        if len(self._min_tracker) == 0:
            self._min_tracker.append(new_node)
        else:
            self._min_tracker.append(min(self._min_tracker[-1], new_node))

        super().push(new_node)

    def pop(self) -> int:
        self._min_tracker.pop()
        return super().pop()

    def min(self) -> int:
        return self._min_tracker[-1]


"""
3.3 Stack of Plates:

Imagine a (literal) stack of plates. If the stack gets too high, it might topple.
Therefore, in real life, we would likely start a new stack when the previous stack exceeds some
threshold. Implement a data structure SetOfStacks that mimics this. SetOfStacks should be
composed of several stacks and should create a new stack once the previous one exceeds capacity.
SetOfStacks. push() and SetOfStacks. pop() should behave identically to a single stack
(that is, pop () should return the same values as it would if there were just a single stack).
FOLLOW UP
Implement a function popAt(int index) which performs a pop operation on a specific substack.
"""


class StackOverFlow(Exception):
    pass


class PopFromEmptyStack(Exception):
    pass


class LimitedStack(MyStack):
    def __init__(self, max_size: int) -> None:
        super().__init__()
        self.max_size = max_size

    def push(self, item: int) -> None:
        if len(self) >= self.max_size:
            raise StackOverFlow
        else:
            super().push(item)

    def pop(self):
        if len(self) <= 0:
            raise PopFromEmptyStack()
        else:
            return super().pop()


class SetOfStacks:
    def __init__(self, substack_capacity=10):
        self.substack_cap = substack_capacity
        self._stacks = []

    def push(self, new_node: int) -> None:
        if len(self._stacks) == 0:
            self._stacks.append(LimitedStack(self.substack_cap))

        try:
            self._stacks[-1].push(new_node)
        except StackOverFlow:
            self._stacks.append(LimitedStack(self.substack_cap))
            self.push(new_node)

    def pop(self) -> int:
        if len(self._stacks) == 0:
            raise IndexError
        try:
            return self._stacks[-1].pop()
        except PopFromEmptyStack:
            self._stacks.pop()
            return self.pop()

    def popAt(self, index: int) -> int:
        return self._stacks[index].pop()


""" 
Queue via Stacks: Implement a MyQueue class which implements a queue using two stacks.

"""


class QueueViaStack:
    def __init__(self) -> None:
        self.forward = MyStack()
        self.backward = MyStack()

    def __len__(self) -> int:
        return max(len(self.forward), len(self.backward))

    def add(self, item: int) -> None:
        if len(self.backward) == 0:
            self.forward.push(item)
        else:
            while len(self.backward) > 0:
                self.forward.push(self.backward.pop())
            self.forward.push(item)

    def remove(self) -> int:
        if len(self.forward) == 0:
            return self.backward.pop()
        else:
            while len(self.forward) > 0:
                self.backward.push(self.forward.pop())
            return self.backward.pop()

    def peek(self) -> int:
        if len(self.forward) == 0:
            return self.backward.peek()
        else:
            while len(self.forward) > 0:
                self.backward.push(self.forward.pop())
            return self.backward.peek()


""" 
3.5 Sort Stack:

Write a program to sort a stack such that the smallest items are on the top. You can use
an additional temporary stack, but you may not copy the elements into any other data structure
(such as an array). The stack supports the following operations: push, pop, peek, and is Empty. 
"""


class SortStack:
    def __init__(self) -> None:
        self.body = MyStack()
        self.temporary = MyStack()

    def __len__(self) -> int:
        return len(self.body)

    def _is_not_ascending(self, item) -> bool:
        if len(self.body) == 0:
            return False
        if item > self.body.peek():
            return True
        else:
            return False

    def push(self, item: int) -> None:
        if self._is_not_ascending(item):
            self.temporary.push(self.body.pop())
        self.body.push(item)
        while len(self.temporary) > 0:
            self.body.push(self.temporary.pop())

    def pop(self) -> int:
        return self.body.pop()

    def peek(self) -> int:
        return self.body.peek()


""" 
3.6 Animal Shelter:

An animal shelter, which holds only dogs and cats, operates on a strictly"first in, first
out" basis. People must adopt either the "oldest" (based on arrival time) of all animals at the shelter,
or they can select whether they would prefer a dog or a cat (and will receive the oldest animal of
that type). They cannot select which specific animal they would like. Create the data structures to
maintain this system and implement operations such as enqueue, dequeueAny, dequeueDog,
and dequeueCat. You may use the built in Linkedlist data structure. 
"""


class Cat:
    pass


class Dog:
    pass


class AnimalShelter:
    def __init__(self) -> None:
        self.body = MyQueue()

    def __len__(self) -> int:
        return len(self.body)

    def enqueue(self, item: Union[Cat, Dog]) -> None:
        self.body.add(item)

    def dequeueAny(self) -> Union[Cat, Dog]:
        return self.body.remove()

    def dequeueDog(self) -> Union[Cat, Dog, None]:
        if type(self.body.peek()) is Dog:
            return self.dequeueAny()
        else:
            return None

    def dequeueCat(self) -> Union[Cat, Dog, None]:
        if type(self.body.peek()) is Cat:
            return self.dequeueAny()
        else:
            return None
