from queue import deque

dx = [1, -1, 0, 0]
dy = [0, 0, 1, -1]

def rotate(x, y, tmp_array, array, step_size):
    for i in range(step_size):
        tmp_array[y+i][x+step_size-1] = array[y][x+i] # case1: (y, ) -> (, x+step)
        tmp_array[y+step_size-1][x+step_size-i-1] = array[y+i][x+step_size-1] # case2: (, x+step) -> (y+step, )
        tmp_array[y+i][x] = array[y+step_size-1][x+i] # case3: (y+step, ) -> (, x)
        tmp_array[y][x+step_size-i-1] = array[y+i][x] # case4: (, x) -> (y, )

    if step_size > 2:
        rotate(x+1, y+1, tmp_array, array, step_size-2)

def ice2water(x, y, adjust_list, array, length):
    cnt = 0
    for i in range(4):
        nx = x + dx[i]
        ny = y + dy[i]
        if 0 <= nx < length and 0 <= ny < length and array[ny][nx] > 0:
            cnt += 1

    if cnt < 3:
        adjust_list.append((x, y))

def BFS(q, visited, array, length):
    global result
    cnt = 1
    while q:
        x, y = q.popleft()
        for i in range(4):
            nx = x + dx[i]
            ny = y + dy[i]
            if 0 <= nx < length and 0 <= ny < length:
                val = array[ny][nx]
                if val > 0 and not visited[ny][nx]:
                    result += val
                    visited[ny][nx] = True
                    cnt += 1
                    q.append((nx, ny))
    return cnt

if __name__ == '__main__':
    n, q = list(map(int, input().split()))
    length = 2 ** n
    array = [list(map(int, input().split())) for _ in range(length)]
    L = list(map(int, input().split()))

    for l in L:
        # rotate for each cell
        step_size = 2 ** l
        if step_size != 1: # rotate condition
            tmp_array = [[0] * length for _ in range(length)]
            for y in range(0, length, step_size):
                for x in range(0, length, step_size):
                    rotate(x, y, tmp_array, array, step_size)
            array = tmp_array
        
        # melt for each cell
        adjust_list = []
        for i in range(length):
            for j in range(length):
                if array[j][i] > 0:
                    ice2water(i, j, adjust_list, array, length)
        for x, y in adjust_list:
            array[y][x] -= 1

    # sum ice amount and find max_block size
    visited = [[False] * length for _ in range(length)]
    max_block = 0
    result = 0 # total ice amount
    q = deque()
    for i in range(length):
        for j in range(length):
            val = array[j][i]
            if val > 0 and not visited[j][i]:
                result += val
                visited[j][i] = True
                q.append((i, j))
                max_block = max(max_block, BFS(q, visited, array, length))
    print(result)
    print(max_block)