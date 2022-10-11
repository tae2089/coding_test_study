import sys
 
input = sys.stdin.readline

n,m = map(int,input().split())
num_list = list(map(int,input().split()))
sum_list = [0]

temp = 0
for i in num_list:
    temp += i
    sum_list.append(temp)

for i in range(m):
    a,b = map(int,input().split())
    print(sum_list[b]-sum_list[a-1])