# Write your MySQL query statement below
select day,emp_id,sum(t) as total_time from 
(select emp_id, event_day as day, out_time-in_time as t from Employees) e 
group by emp_id, day;
