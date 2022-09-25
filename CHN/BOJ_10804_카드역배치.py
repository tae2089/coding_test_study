a = list(range(21))
for _ in range(10) :
  start, fin = map(int, input().split())
  for i in range((fin-start+1)//2):
    a[start+i], a[fin-i] = a[fin-i], a[start+i]

a.pop(0)  
for x in a:
  print(x, end = ' ')
