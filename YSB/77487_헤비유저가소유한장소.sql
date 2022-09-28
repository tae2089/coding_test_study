https://school.programmers.co.kr/learn/courses/30/lessons/77487

select b.*
from places as b
where b.host_id in 
(select a.host_id
from (select host_id, count(*) as cnt 
      from places 
      group by host_id) as a
where a.cnt >=2)
order by id
