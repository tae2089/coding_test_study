from collections import defaultdict
import sys
input = sys.stdin.readline

def dfs(dic, visited, start_point):
    result = 1
    for i in dic[start_point]:
        if not visited[i]:
            visited[i] = True
            result += dfs(dic, visited, i)
    return result

if __name__ == '__main__':
    n, m = map(int, input().split())

    heavy_dict = defaultdict(lambda: list())
    light_dict = defaultdict(lambda: list())
    for _ in range(m):
        heavy, light = map(int, input().split())
        heavy_dict[light-1].append(heavy-1)
        light_dict[heavy-1].append(light-1)
    
    ans = 0
    for i in range(n):
        visited = [False] * n
        visited[i] = True
        result = dfs(heavy_dict, visited, i) - 1

        if result >= (n+1)//2:
            ans += 1
            continue
        
        visited = [False] * n
        visited[i] = True
        result = dfs(light_dict, visited, i) - 1

        if result >= (n+1)//2:
            ans += 1
            continue
    print(ans)