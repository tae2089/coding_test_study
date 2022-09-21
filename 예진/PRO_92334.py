def solution(id_list, report, k):
    report = list(set(report)) # 신고 중복제거
    reported_ct = {x : 0 for x in id_list} # 유저별 신고 당한 횟수
    ans_dict = {x : 0 for x in id_list} # 유저별 받는 처리 결과 메일 수
    
    for r in report: # 유저별 신고 당한 횟수 count
        reported_ct[r.split()[1]] += 1
    
    for r in report:
        a, b = r.split()
        if reported_ct[b] >= k: # 신고당한 횟수가 k보다 크면
            ans_dict[a] += 1 # 유저별 수신 메일 수 +1
        
    return list(ans_dict.values())