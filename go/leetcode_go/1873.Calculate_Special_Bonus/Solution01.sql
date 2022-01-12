# Write your MySQL query statement below

select employee_id, if(substr(name, 1,1) != 'M' and employee_id%2=1, salary, 0) as bonus 
from Employees 
order by employee_id asc;
