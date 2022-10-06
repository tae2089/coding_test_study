# 열린 괄호를 만나면 -> 스택에 넣어주고
# 닫힌 괄호를 만나면 -> stack top을 꺼내 쌍을 맞춰준다. 이때 쌍이 맞지 않으면 잘못된 괄호 입력.



b = list(input())

stack =[]
answer = 0
tmp =1

# 인덱스로 돌아보자
for i in range(len(b)):
  if b[i] =='(':
    stack.append(b[i])
    tmp *=2
  
  elif b[i] =='[':
    stack.append(b[i])
    tmp *=3
  
  elif b[i] ==')':
    if not stack or stack[-1] =='[': 
      answer = 0
      break
    if b[i-1] =="(":
      answer += tmp
    stack.pop() #  
    tmp //=2 ## 중복을 나눠주는 느낌

  else:
    if not stack or stack[-1] =="(":
      answer =0
      break
    if b[i-1] =='[':
      answer +=tmp
    stack.pop()
    tmp //= 3

if stack:
  print(0)
else:
  print(answer)
