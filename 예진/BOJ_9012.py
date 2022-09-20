for _ in range(int(input())):
    x = input()
    while True:
        x = x.replace('()', '')
        if x.replace('()', '') == x: 
            break
    if len(x) : print('NO')
    else: print("YES")