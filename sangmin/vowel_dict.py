from itertools import product
def solution(word):
    dict = []
    for i in range(1, 6):
        for v in product(['A','E','I','O','U'], repeat=i):
            dict.append(''.join(list(v)))
    dict.sort()
    
    return dict.index(word) + 1