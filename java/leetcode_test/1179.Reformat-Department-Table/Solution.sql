# Write your MySQL query statement below

# 1
select id,
    sum(case month when 'Jan' then revenue end) as Jan_revenue,
    sum(case month when 'Feb' then revenue end) as Feb_revenue,
    sum(case month when 'Mar' then revenue end) as Mar_revenue,
    sum(case month when 'Apr' then revenue end) as Apr_revenue,
    sum(case month when 'May' then revenue end) as May_revenue,
    sum(case month when 'Jun' then revenue end) as Jun_revenue,
    sum(case month when 'Jul' then revenue end) as Jul_revenue,
    sum(case month when 'Aug' then revenue end) as Aug_revenue,
    sum(case month when 'Sep' then revenue end) as Sep_revenue,
    sum(case month when 'Oct' then revenue end) as Oct_revenue,
    sum(case month when 'Nov' then revenue end) as Nov_revenue,
    sum(case month when 'Dec' then revenue end) as Dec_revenue
from Department group by id;