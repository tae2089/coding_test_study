import sys
from queue import deque
input = sys.stdin.readline

if __name__ == '__main__':
    n = int(input())
    candidate_list = []
    
    for i in range(n):
        candidate_list.append(input().strip().split())
        
    q = deque(input().strip().split())
    while q:
        curr = q.popleft()
        done = 0
        for idx in range(n):
            if len(candidate_list[idx]) > 0 and candidate_list[idx][0] == curr:
                candidate_list[idx] = candidate_list[idx][1:]
                done = 1
                break
        if done == 0:
            q.append(curr)
            break
            
    if q or sum([len(i) for i in candidate_list]) > 0:
        print('Impossible')
    else:
        print('Possible')