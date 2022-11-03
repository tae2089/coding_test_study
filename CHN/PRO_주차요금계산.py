def solution(fees, records):
  minute, won, unitM, unitWon = fees

  park = dict()
  answer = defaultdict(int)

  for cmm in records: 
    time, number, order = cmm.split()
    h, m = map(int ,time.split(':'))
    toMin = h*60 + m 

    if order == "IN" : 
      park[number] = toMin 
    else : 
      delay = toMin - park[number]
      answer[number] +=delay
      del park[number]

  for number in park : 
    delay = 23*60+59 - park[number]
    answer[number] += delay
  
  wonList = []
  keys = sorted(answer.keys()) 
  
  for number in keys: 
    addTime = ceil((answer[number]- minute)/unitM)
    addfee = addTime *unitWon
    wonList.append(won+max(0, addfee))

  return wonList
