# Write your MySQL query statement below

# 1
/*
select class from (
    select class, student from courses group by class, student
) as c 
group by class having count(class) >= 5; 
*/

# 2
/*
select class from (
    select distinct student, class from courses 
) as c group by class having count(class) >= 5;
*/

# 3
select class from courses group by class having count(distinct(student)) >= 5;