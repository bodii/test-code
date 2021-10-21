select d.name as Department,e.name as Employee,e.Salary
from Employee as e join
Department as d ON e.DepartmentId = d.Id
where
(e.DepartmentId , e.Salary) 
in
(select DepartmentId, MAX(Salary) as Salary from Employee group by DepartmentId);