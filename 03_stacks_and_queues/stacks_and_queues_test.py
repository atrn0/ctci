import unittest

from stacks_and_queues import *


class TestMyStack(unittest.TestCase):
    def test_methods(self):
        s = MyStack()
        self.assertEqual(len(s), 0)
        s.push(3)
        self.assertEqual(len(s), 1)
        self.assertEqual(s.peek(), 3)
        s.push(34)
        self.assertEqual(len(s), 2)
        self.assertEqual(s.peek(), 34)
        self.assertEqual(s.pop(), 34)
        self.assertEqual(len(s), 1)
        s.pop()
        self.assertEqual(len(s), 0)
        with self.assertRaises(IndexError):
            s.peek()
        with self.assertRaises(IndexError):
            s.pop()


class TestMyQueue(unittest.TestCase):
    def test_methods(self):
        q = MyQueue()
        self.assertEqual(len(q), 0)
        q.add(3)
        self.assertEqual(len(q), 1)
        self.assertEqual(q.peek(), 3)
        q.add(34)
        self.assertEqual(len(q), 2)
        self.assertEqual(q.peek(), 3)
        self.assertEqual(q.remove(), 3)
        self.assertEqual(len(q), 1)
        q.remove()
        self.assertEqual(len(q), 0)
        with self.assertRaises(IndexError):
            q.peek()
        with self.assertRaises(IndexError):
            q.remove()


class TestLinkedListStack(unittest.TestCase):
    def test_linkedliststack_methods(self):
        s = LinkedListStack()
        self.assertEqual(len(s), 0)
        s.push(3)
        self.assertEqual(len(s), 1)
        self.assertEqual(s.peek(), 3)
        s.push(34)
        self.assertEqual(len(s), 2)
        self.assertEqual(s.peek(), 34)
        self.assertEqual(s.pop(), 34)
        self.assertEqual(len(s), 1)
        s.pop()
        self.assertEqual(len(s), 0)
        with self.assertRaises(IndexError):
            s.peek()
        with self.assertRaises(IndexError):
            s.pop()


class TestStackMin(unittest.TestCase):
    def test_methods(self):
        s = StackMin()
        self.assertEqual(len(s), 0)
        s.push(2)
        self.assertEqual(len(s), 1)
        self.assertEqual(s.peek(), 2)
        self.assertEqual(s.min(), 2)
        s.push(1)
        self.assertEqual(len(s), 2)
        self.assertEqual(s.peek(), 1)
        self.assertEqual(s.min(), 1)

        self.assertEqual(s.pop(), 1)
        self.assertEqual(s.min(), 2)


class TestLimitedStack(unittest.TestCase):
    def test_pop_from_empty(self):
        s = LimitedStack(3)
        self.assertEqual(len(s), 0)
        with self.assertRaises(PopFromEmptyStack):
            s.pop()

    def test_stack_capacity(self):
        s = LimitedStack(3)
        s.push(1)
        s.push(2)
        s.push(3)
        self.assertEqual(len(s), 3)
        with self.assertRaises(StackOverFlow):
            s.push(4)


class TestSetOfStacks(unittest.TestCase):
    def test_methods(self):
        s = SetOfStacks(substack_capacity=2)
        with self.assertRaises(IndexError):
            s.pop()
        s.push(1)
        s.push(2)
        self.assertEqual(len(s._stacks), 1)
        s.push(3)
        self.assertEqual(len(s._stacks), 2)
        self.assertEqual(s.pop(), 3)
        self.assertEqual(len(s._stacks), 2)
        self.assertEqual(s.pop(), 2)
        self.assertEqual(len(s._stacks), 1)


class TestQueueViaStack(unittest.TestCase):
    def test_methods(self):
        q = MyQueue()
        self.assertEqual(len(q), 0)
        q.add(3)
        self.assertEqual(len(q), 1)
        self.assertEqual(q.peek(), 3)
        q.add(34)
        self.assertEqual(len(q), 2)
        self.assertEqual(q.peek(), 3)
        self.assertEqual(q.remove(), 3)
        self.assertEqual(len(q), 1)
        q.remove()
        self.assertEqual(len(q), 0)
        with self.assertRaises(IndexError):
            q.peek()
        with self.assertRaises(IndexError):
            q.remove()


class TestSortStack(unittest.TestCase):
    def test_methods(self):
        s = SortStack()
        self.assertEqual(len(s), 0)
        s.push(3)
        self.assertEqual(len(s), 1)
        self.assertEqual(s.peek(), 3)
        s.push(4)
        self.assertEqual(s.peek(), 3)
        s.push(2)
        self.assertEqual(s.peek(), 2)
        self.assertEqual(s.pop(), 2)
        self.assertEqual(s.peek(), 3)
        self.assertEqual(s.pop(), 3)


class TestAnimalShelter(unittest.TestCase):
    def test_enqueue(self):
        q = AnimalShelter()
        self.assertEqual(len(q), 0)
        q.enqueue(Dog())
        q.enqueue(Cat())
        q.enqueue(Cat())
        q.enqueue(Dog())
        self.assertEqual(len(q), 4)
        return q

    def test_dequeany(self):
        q = self.test_enqueue()
        self.assertEqual(type(q.dequeueAny()), type(Dog()))

    def test_dequeDog(self):
        q = self.test_enqueue()
        self.assertEqual(type(q.dequeueDog()), type(Dog()))
        self.assertEqual(type(q.dequeueDog()), type(None))

    def test_dequeCat(self):
        q = self.test_enqueue()
        self.assertEqual(type(q.dequeueCat()), type(None))
        q.dequeueAny()
        self.assertEqual(type(q.dequeueCat()), type(Cat()))


if __name__ == "__main__":
    unittest.main()
