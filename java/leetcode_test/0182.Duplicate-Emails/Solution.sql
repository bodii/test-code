# Write your MySQL query statement below

/*
# 1
select Email
from (select Email, count(Email) e
    from Person
    group by Email
    having e > 1) as p;
*/

/*
# 2
select Email
from (select Email, count(Email) e
    from Person
    group by Email) as p where e > 1;
*/

# 3
select Email
from Person
group by Email
having count(Email) > 1;