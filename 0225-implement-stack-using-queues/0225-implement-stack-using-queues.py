class MyStack:

    def __init__(self):
        self.myStack = []
        self.myStackLen = 0

    def push(self, x: int) -> None:
        self.myStack.append(x)
        self.myStackLen += 1

    def pop(self) -> int:
        removed_number = self.myStack[self.myStackLen-1]
        self.myStack.pop(self.myStackLen-1)
        self.myStackLen -= 1
        return removed_number

    def top(self) -> int:
        return self.myStack[self.myStackLen-1]

    def empty(self) -> bool:
        if self.myStackLen == 0:
            return True
        return False


# Your MyStack object will be instantiated and called as such:
# obj = MyStack()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.top()
# param_4 = obj.empty()