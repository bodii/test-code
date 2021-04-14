SELECT d.Name AS Department,
		 e.Name AS Employee,
		 e.Salary
FROM Department d, Employee e
WHERE DepartmentId = d.Id
		AND 
    (SELECT count(distinct Salary)
    FROM Employee
    WHERE DepartmentId=d.Id
    		AND Salary > e.Salary )<3
ORDER BY  DepartmentId asc,e.Salary DESC, e.Id desc;