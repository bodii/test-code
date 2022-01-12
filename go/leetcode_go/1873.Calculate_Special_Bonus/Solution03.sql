select
employee_id,
salary * (employee_id&1 and left(name, 1)<>'M') as bonus
from employees
order by employee_id;

