SET @hour := -1; -- 변수 선언

SELECT 
    (@hour := @hour + 1) as HOUR, -- 0부터 하나씩 증가
    (SELECT count(*) FROM ANIMAL_OUTS WHERE hour(DATETIME) = @hour) 
    as COUNT -- 변수랑 DATETIME의 hour 부분이 같은 수 세기
FROM ANIMAL_OUTS
WHERE @hour < 23; -- 23까지