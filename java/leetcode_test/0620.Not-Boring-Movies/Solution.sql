# Write your MySQL query statement below

# 1
/*
select id, movie, description, rating from cinema 
where description <> 'boring' and id % 2 = 1 order by rating desc;
*/

# 2
select id, movie, description, rating from cinema 
where mod(id, 2) = 1 and description <> 'boring' order by rating desc;