select v1.Department, v1.Employee, v1.Salary from 
(select d.Name as Department, e.Name as Employee, e.Salary as Salary from Employee as e left join Department as d on e.DepartmentId = d.Id) as v1,
(select d.Name as Department, e.Name as Employee, max(e.Salary) as Salary from Employee as e left join Department as d on e.DepartmentId = d.Id
group by e.DepartmentId order by e.Salary) as v2
where v1.Department=v2.Department and v1.Salary=v2.Salary;