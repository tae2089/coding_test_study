
from collections import deque

n, m = map(int, input().split())
board = [] # graph
for i in range(n):
  board.append(list(input())) 
  
dx = [0,0,1,-1]
dy = [1,-1,0,0]

# 4차원 visited 배열 
# 사용이유 : 빨간색 구슬의 좌표는 이전과 같을지라도 파란색 구슬의 좌표가 다르다면 다른 상황이기 때문,, 이라고 한다.
visited = [[[[False]* m for _ in range(n)] for _ in range(m)]for _ in range(n)]


def roll(x, y, direct): # 시뮬레이션,,
  cnt = 0
  # while문을 써서 기울이는 이동 구현
  while board[x+dx[direct]][y+dy[direct]] != '#' and board[x][y] != 'O': # 벽,장애물,출구가 아니면
    x +=dx[direct] # 이동
    y +=dy[direct] # 이동
    cnt += 1

  return x,y,cnt # 최종 이동 좌표, 이동 횟수 반환



def bfs(ra, rb, ba, bb): # r = red, b= blue
  queue = deque()
  queue.append((ra, rb, ba, bb, 1)) 

  while queue:
    rx, ry, bx, by, cnt = queue.popleft()
    visited[rx][ry][bx][by] = True
    
    if cnt>10:
      print(-1)
      exit(0) # exit()이란? 코드 강제 종료하고 싶을 때 사용 exit(0) : 성공적으로 프로그램 종료 (EXIT_SUCCESS)/exit(1) : 성공적으로 프로그램 종료X (EXIT_FAILURE)
    
    for i in range(4): # dx, dy 테크닉
      nrx, nry, rcnt = roll(rx, ry,i)
      nbx, nby, bcnt = roll(bx, by,i)

      if board[nbx][nby] != 'O':
        if board[nrx][nry] == 'O':
          print(cnt)
          exit(0)

        # 이동하다가 겹쳤을때, 앞뒤 구분해 재배치
        if nrx == nbx and nry == nby:
          if rcnt >bcnt :
            nrx -= dx[i]
            nby -= dy[i]
          else:
            nbx -= dx[i]
            nby -= dy[i]
        
        if not visited[nrx][nry][nbx][nby]:
          visited[nrx][nry][nbx][nby] = True
          queue.append((nrx, nry,nbx, nby, cnt+1))

  print(-1)
  return

rx, ry, bx, by = 0,0,0,0
for i in range(n):
  for j in range(m):
    if graph[i][j] =='R':
      rx, ry = i, j
    if graph[i][j]=='B':
      bx, by = i, j

bfs(rx, ry, bx, by)
