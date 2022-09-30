a = int(input())

def func(x):
    n = str(x)
    dif = int(n[1]) - int(n[0])
    for i in range(len(n)-1):
        if int(n[i+1])-int(n[i]) != dif:
            return 0
    return 1

ans = 9
if a <10:
    print(a)
    
else:
    for j in range(10,a+1):
        ans += func(j)

    print(ans)

    