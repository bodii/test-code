# Write your MySQL query statement below

# 1
/*
update salary set sex = case when sex='f' then 'm' else 'f' end;
*/

# 2
/*
update salary set sex = if(sex='f','m','f');
*/

# 3
update salary set sex = case sex when 'f' then 'm' else 'f' end;