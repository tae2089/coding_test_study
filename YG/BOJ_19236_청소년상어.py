dx = [0, -1, -1, -1, 0, 1, 1, 1]
dy = [-1, -1, 0, 1, 1, 1, 0, -1]

def move_fishes(fish, space, fish_list):
    for i in fish_list:
        y, x, direction = fish[i]
        for j in range(8):
            d = (direction + j) % 8
            nx = x + dx[d]
            ny = y + dy[d]
            if 0 <= nx < 4 and 0 <= ny < 4 and space[ny][nx] != 17:
                if space[ny][nx] != 0:
                    fish[i] = (ny, nx, d)
                    fish[space[ny][nx]] = (y, x, fish[space[ny][nx]][2])

                    space[y][x] = space[ny][nx]
                    space[ny][nx] = i
                else:
                    fish[i] = (ny, nx, d)
                    space[ny][nx] = i
                    space[y][x] = 0
                break

def shark(y, x, d, fish, space, fish_list):
    ans_list = [0]
    for i in range(1, 4):
        nx = x + dx[d] * i
        ny = y + dy[d] * i
        if 0 <= nx < 4 and 0 <= ny < 4 and 0 < space[ny][nx] < 17:
            f = dict(fish)
            s = [i[:] for i in space]
            fl = fish_list[:]

            y_, x_, d_ = f[s[ny][nx]]
            fl.remove(s[ny][nx])
            ans = s[y_][x_]
            s[y][x] = 0
            s[y_][x_] = 17
            move_fishes(f, s, fl)
            ans_list.append(ans + shark(y_, x_, d_, f, s, fl))

    return max(ans_list)

if __name__ == '__main__':
    space = [[0] * 4 for _ in range(4)]
    fish_list = [i for i in range(1, 17)]
    fish = dict()
    for i in range(4):
        fish_num_direction = input().split()
        for j in range(4):
            fish[int(fish_num_direction[2*j])] = (i, j, int(fish_num_direction[2*j+1]) - 1)
            space[i][j] = int(fish_num_direction[2*j])

    y, x, d = fish[space[0][0]]
    fish_list.remove(space[0][0])
    ans = space[0][0]
    space[0][0] = 17
    move_fishes(fish, space, fish_list)

    print(ans + shark(y, x, d, fish, space, fish_list))