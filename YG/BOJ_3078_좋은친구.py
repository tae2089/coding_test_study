import sys
from queue import deque
from collections import defaultdict
input = sys.stdin.readline

if __name__ == '__main__':
    n, k = map(int, input().strip().split())
    q = deque()
    candidate = defaultdict(lambda: 0)
    ans = 0
    for i in range(n):
        curr = len(input().strip())
        if len(q) == k:
            ans += candidate[curr]
            tmp = q.popleft()
            candidate[tmp] -= 1
            q.append(curr)
            candidate[curr] += 1
        else:
            ans += candidate[curr]
            q.append(curr)
            candidate[curr] += 1
    print(ans)