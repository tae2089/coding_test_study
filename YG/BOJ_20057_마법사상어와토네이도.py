dx = [-1, 0, 1, 0]
dy = [0, 1, 0, -1]

def sand_move(x, y, amount, array, direction, ans, n):
    total = 0

    target = int(amount * 0.01)
    for i in [1, 3]:
        nx = x + dx[(direction + i) % 4]
        ny = y + dy[(direction + i) % 4]
        if 0 <= nx < n and 0 <= ny < n: # move (inside sand)
            array[ny][nx] += target
        else: # remove (outside sand)
            ans += target
        total += target

    target = int(amount * 0.07)
    for i in [1, 3]:
        nx = x + dx[direction] + dx[(direction + i) % 4]
        ny = y + dy[direction] + dy[(direction + i) % 4]
        if 0 <= nx < n and 0 <= ny < n: # move (inside sand)
            array[ny][nx] += target
        else: # remove (outside sand)
            ans += target
        total += target

    target = int(amount * 0.02)
    for i in [1, 3]:
        nx = x + dx[direction] + dx[(direction + i) % 4] * 2
        ny = y + dy[direction] + dy[(direction + i) % 4] * 2
        if 0 <= nx < n and 0 <= ny < n: # move (inside sand)
            array[ny][nx] += target
        else: # remove (outside sand)
            ans += target
        total += target

    target = int(amount * 0.1)
    for i in [1, 3]:
        nx = x + dx[direction] * 2 + dx[(direction + i) % 4]
        ny = y + dy[direction] * 2 + dy[(direction + i) % 4]
        if 0 <= nx < n and 0 <= ny < n: # move (inside sand)
            array[ny][nx] += target
        else: # remove (outside sand)
            ans += target
        total += target

    target = int(amount * 0.05)
    nx = x + dx[direction] * 3
    ny = y + dy[direction] * 3
    if 0 <= nx < n and 0 <= ny < n: # move (inside sand)
        array[ny][nx] += target
    else: # remove (outside sand)
        ans += target
    total += target

    target = amount - total
    nx = x + dx[direction] * 2
    ny = y + dy[direction] * 2
    if 0 <= nx < n and 0 <= ny < n: # move (inside sand)
        array[ny][nx] += target
    else: # remove (outside sand)
        ans += target

    return ans

def move(x, y, array, step_size, direction, ans, n):
    for _ in range(step_size):
        # target point
        nx = x + dx[direction]
        ny = y + dy[direction] 

        # target point's sand move
        if array[ny][nx] != 0:
            amount = array[ny][nx]
            array[ny][nx] = 0
            ans = sand_move(x, y, amount, array, direction, ans, n)
        x, y = nx, ny # move to target point
    return x, y, ans

if __name__ == '__main__':
    n = int(input())
    array = [list(map(int, input().split())) for _ in range(n)]
    center = n // 2 # initial point
    direction = 0 # left side
    step_size = 1

    x, y = center, center

    ans = 0
    while (x, y) != (0, 0): # terminal condition
        x, y, ans = move(x, y, array,  step_size, direction, ans, n)
        direction = (direction + 1) % 4
        x, y, ans = move(x, y, array, step_size, direction, ans, n)
        direction = (direction + 1) % 4
        if step_size == n - 1: # terminal condition
            x, y, ans = move(x, y, array, step_size, direction, ans, n)
            break
        step_size += 1
    print(ans)