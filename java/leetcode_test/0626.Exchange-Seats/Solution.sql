# Write your MySQL query statement below

# 1
/*
select 
    (case when mod(id, 2) = 1 and id = count then id
    else id + 1
    end) as id, student 
from seat, 
    (select count(1) as count from seat) as c
order by id asc;
*/

# 2
select (id - 1) as id, student from seat where mod(id, 2) = 0
union
select (case when id=max_id then id else id + 1 end) as id, student from seat,(select max(id) as max_id from seat) as c where mod(id, 2) = 1
order by id asc;