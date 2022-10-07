import heapq
import sys
input = sys.stdin.readline

if __name__ == '__main__':
    n = int(input())
    min_heap = []
    max_heap = []
    for i in range(n):
        target_num = int(input())
        if i == 0:
            heapq.heappush(min_heap, -target_num)
            print(target_num)
        else:
            if (i + 1) % 2 == 0:
                heapq.heappush(max_heap, target_num)
            else:
                heapq.heappush(min_heap, -target_num)
            num1 = -heapq.heappop(min_heap)
            num2 = heapq.heappop(max_heap)
            if num1 < num2:
                heapq.heappush(min_heap, -num1)
                heapq.heappush(max_heap, num2)
            else:
                heapq.heappush(min_heap, -num2)
                heapq.heappush(max_heap, num1)
            print(-min_heap[0])