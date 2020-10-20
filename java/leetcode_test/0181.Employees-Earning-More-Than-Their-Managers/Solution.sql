# Write your MySQL query statement below
# 1
select e1.Name as Employee
from Employee e1, Employee e2
where e1.ManagerId=e2.Id and e1.Salary > e2.Salary;

# 2
select e1.Name as Employee
from Employee e1 inner join Employee e2
on e1.ManagerId=e2.Id where e1.Salary > e2.Salary;
