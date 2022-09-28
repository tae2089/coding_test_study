from collections import defaultdict
dx = [0, 1, 1, 1, 0, -1, -1, -1]
dy = [-1, -1, 0, 1, 1, 1, 0, -1]

def check(direction):
    if direction[0] in [0, 2, 4, 6]: # even number
        for d in direction:
            if d in [1, 3, 5, 7]: # odd number
                return [1, 3, 5, 7]
    else: # odd number
        for d in direction:
            if d in [0, 2, 4, 6]: # even number
                return [1, 3, 5, 7]
    return [0, 2, 4, 6]

def move_fireball(matrix_to_fireball, fireball_dict, n):
    tmp_matrix_to_fireball = defaultdict(lambda: list()) # this dictionary includes that which fireball is in each (row, column)
    for k, v in matrix_to_fireball.items():
        for v_ in v:
            mass, velocity, direction = fireball_dict[v_]
            ny = k[0] + dy[direction] * velocity
            nx = k[1] + dx[direction] * velocity
            # if fireball crosses the bounds of the matrix, it goes back to the opposite end of the matrix
            if nx < 0 or nx >= n:
                nx = nx % n
            if ny < 0 or ny >= n:
                ny = ny % n
            tmp_matrix_to_fireball[(ny, nx)].append(v_)
    return tmp_matrix_to_fireball            

def merge_fireball(matrix_to_fireball, fireball_dict, cnt):
    for k, v in matrix_to_fireball.items():
        if len(v) >= 2: # merge condition
            total_mass = 0
            total_velocity = 0
            total_direction = []
            for v_ in v:
                mass, velocity, direction = fireball_dict.pop(v_)
                total_mass += mass
                total_velocity += velocity
                total_direction.append(direction)
            each_mass = total_mass // 5
            each_velocity = total_velocity // len(total_direction)
            matrix_to_fireball[k] = [] # corresponding (row, column) initialization
            if each_mass > 0:
                for d in check(total_direction):
                    matrix_to_fireball[k].append(cnt)
                    fireball_dict[cnt] = (each_mass, each_velocity, d)
                    cnt += 1
    return cnt

if __name__ == '__main__':
    n, m, k = list(map(int, input().split()))
    matrix_to_fireball = defaultdict(lambda: list()) # this dictionary includes that which fireball is in each (row, column)
    fireball_dict = dict() # this dictionary includes the mass, velocity, direction of each fireball
    # initialization
    for i in range(m):
        r, c, mass, velocity, direction = list(map(int, input().split()))
        matrix_to_fireball[(r-1, c-1)].append(i)
        fireball_dict[i] = (mass, velocity, direction)
    cnt = m + 1 # total fireball number variable
    for i in range(k):
        matrix_to_fireball = move_fireball(matrix_to_fireball, fireball_dict, n)
        cnt = merge_fireball(matrix_to_fireball, fireball_dict, cnt)
    
    print(sum([v[0] for v in fireball_dict.values()])) # mass summation