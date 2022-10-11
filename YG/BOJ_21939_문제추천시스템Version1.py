from collections import defaultdict
import heapq
import sys
input = sys.stdin.readline

n = int(input())
min_q = []
max_q = []
problem_list = defaultdict(lambda: 0)
for _ in range(n):
    p, l = map(int, input().split())
    heapq.heappush(min_q, (l, p))
    heapq.heappush(max_q, (-l, -p))
    problem_list[p].add(l)
    
removed_list = defaultdict(lambda: set())
m = int(input())
for _ in range(m):
    command = input().split()
    if command[0] == 'recommend':
        if int(command[1]) == 1:
            while True:
                if -max_q[0][0] in removed_list[-max_q[0][1]]:
                    heapq.heappop(max_q)
                else:
                    break
            print(-max_q[0][1])
        else:
            while True:
                if min_q[0][0] in removed_list[min_q[0][1]]:
                    heapq.heappop(min_q)
                else:
                    break
            print(min_q[0][1])
    elif command[0] == 'add':
        difficulty = int(command[2])
        problem = int(command[1])
        heapq.heappush(min_q, (difficulty, problem))
        heapq.heappush(max_q, (-difficulty, -problem))
        problem_list[problem].add(difficulty)
    else: # solved
        removed_list[int(command[1])] = removed_list[int(command[1])].union(problem_list[int(command[1])])