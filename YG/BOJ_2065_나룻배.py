import sys
from queue import deque
input = sys.stdin.readline

if __name__ == '__main__':
    m, t, n = map(int, input().strip().split())
    time = 0
    curr = 'left'
    left_q = deque()
    right_q = deque()
    for i in range(n):
        arrive_time, location = input().strip().split()
        if location == 'left':
            left_q.append((int(arrive_time), i))
        else:
            right_q.append((int(arrive_time), i))
        
    answer = [0] * (len(left_q) + len(right_q))
    while left_q or right_q:
        cnt = 0
        if curr == 'left':
            if left_q:
                left = left_q[0][0]
            else:
                left = 100001
            if right_q:
                right = right_q[0][0]
            else:
                right = 100001
                
            if left <= time or left <= right: # left start
                time = max(time, left)
                while left_q and left_q[0][0] <= time and cnt < m:
                    _, idx = left_q.popleft()
                    cnt += 1
                    answer[idx] = time + t
                time += t
                curr = 'right'
            else: # right start
                time = max(time, right) + t
                while right_q and right_q[0][0] <= time and cnt < m:
                    _, idx = right_q.popleft()
                    cnt += 1
                    answer[idx] = time + t
                time += t
        else:
            if left_q:
                left = left_q[0][0]
            else:
                left = 100001
            if right_q:
                right = right_q[0][0]
            else:
                right = 100001
                
            if right <= time or right <= left: # right start
                time = max(time, right)
                while right_q and right_q[0][0] <= time and cnt < m:
                    _, idx = right_q.popleft()
                    cnt += 1
                    answer[idx] = time + t
                time += t
                curr = 'left'
            else: # left start
                time = max(time, left) + t
                while left_q and left_q[0][0] <= time and cnt < m:
                    _, idx = left_q.popleft()
                    cnt += 1
                    answer[idx] = time + t
                time += t
            
    for i in answer:
        print(i)