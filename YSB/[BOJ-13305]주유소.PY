n = int(input())
len_street = list(map(int,input().split()))
fee = list(map(int,input().split()))

minfee = fee[0]
sum = 0

for i in range(n-1):
    if minfee > fee[i]:
        minfee = fee[i]
    sum += minfee * len_street[i]

print(sum)
