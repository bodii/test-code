select v1.Department, v1.Employee, v1.Salary from 
(select d.Name as Department, e.Name as Employee, e.Salary as Salary, e.DepartmentId from Employee as e left join Department as d on e.DepartmentId = d.Id) as v1,
(select DepartmentId, max(Salary) as Salary from Employee group by DepartmentId) as v2
where v1.DepartmentId=v2.DepartmentId and v1.Salary=v2.Salary;