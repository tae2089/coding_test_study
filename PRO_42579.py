def solution(genres, plays):
    total = {x: 0 for x in set(genres)}
    for i in range(len(genres)):
        total[genres[i]] += plays[i] # 장르별 재생 횟수 합

    g_total = {}
    for i, (g, p) in enumerate(zip(genres, plays)):
        if g not in g_total:
            g_total[g] = [(i, p)]
        else:
            g_total[g].append((i, p)) # 장르별 재생 횟수 list

    ans = []
    for (k, v) in sorted(total.items(), key = lambda x: x[1], reverse = True): # 재생횟수 높은 장르에서
        for (i, p) in sorted(g_total[k], key = lambda x: x[1], reverse = True)[:2]: # 베스트 재생 2개씩
            ans.append(i) # 정답에 추가

    return ans
