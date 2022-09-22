from queue import deque

def BFS(q, visited, min_dist = 401):
    while q:
        x, y, dist = q.popleft()
        if dist > min_dist:
            continue
        visited[x][y] = dist
        
        for i in range(4):
            nx = x + dx[i]
            ny = y + dy[i]
            if 0 <= nx < n and 0 <= ny < n and not visited[nx][ny]:
                if 0 < space[nx][ny] < size:
                    visited[nx][ny] = dist + 1
                    min_dist = min(min_dist, dist + 1)
                    q.append((nx, ny, dist + 1))
                elif space[nx][ny] == 0 or space[nx][ny] == size:
                    visited[nx][ny] = dist + 1
                    q.append((nx, ny, dist + 1))
    return visited, min_dist

def check(visited, min_dist):
    global answer, size, eat
    for i in range(n):
        for j in range(n):
            if visited[i][j] == min_dist and 0 < space[i][j] < size:
                answer += visited[i][j]
                eat += 1
                if eat == size:
                    size += 1
                    eat = 0
                space[i][j] = 0
                curr_x, curr_y = i, j
                return curr_x, curr_y
    return -1, -1

if __name__ == '__main__':
    n = int(input()) # space size
    space = []
    cnt = 0
    for i in range(n):
        state = list(map(lambda x: int(x), input().split()))
        space.append(state)
        if cnt == 1:
            continue
        if 9 in state:
            curr_x = i
            for j in range(n):
                if state[j] == 9:
                    curr_y = j
                    break

    space[curr_x][curr_y] = 0
    dx = [0, 0, 1, -1]
    dy = [1, -1, 0, 0]
    size = 2
    eat = 0
    answer = 0
    q = deque()
    while True:
        visited = [[False] * n for _ in range(n)]
        q.append((curr_x, curr_y, 0))
        visited, min_dist = BFS(q, visited)
        curr_x, curr_y = check(visited, min_dist)
        if curr_x == -1:
            break
    print(answer)