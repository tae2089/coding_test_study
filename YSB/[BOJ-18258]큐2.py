from collections import deque
import sys
data = deque()
n = int(sys.stdin.readline())

def push(x):
    data.append(x)

def pop():
    if len(data) == 0:
        print(-1)
    else:
        res = data.popleft()
        print(res)

def size():
    print(len(data))

def empty():
    if len(data) == 0:
        print(1)
    else:
        print(0)

def front():
    if len(data) == 0:
        print(-1)
    else:
        print(data[0])
        
def back():
    if len(data) == 0:
        print(-1)
    else:
        print(data[-1])

for _ in range(n):
    ord = sys.stdin.readline().split()
    if ord[0] == 'pop':
        pop()
    elif ord[0] == 'size':
        size()
    elif ord[0] =='empty':
        empty()
    elif ord[0] == 'front':
        front()
    elif ord[0] == 'back':
        back()
    else:
        push(ord[1])
