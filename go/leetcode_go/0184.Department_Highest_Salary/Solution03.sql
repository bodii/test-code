select d.Name as Department, e.Name as Employee, e.Salary from Employee as e 
left join
(select DepartmentId, max(Salary) as Salary from Employee group by DepartmentId) as v 
on v.DepartmentId = e.DepartmentId
left join 
Department as d 
on e.DepartmentId = d.Id 
where e.DepartmentId=v.DepartmentId and e.Salary=v.Salary;