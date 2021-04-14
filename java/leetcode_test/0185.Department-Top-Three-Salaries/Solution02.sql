SELECT Department,
		Employee,
		Salary
FROM 
    (
        SELECT e.Name AS Department,
		 e.Employee,
		 e.Salary,
		 if(@tmp_Salary != e.Salary
    		OR @tmp_Salary is null, @tmp_Salary:=e.Salary, @tmp_Salary:=0) AS tmp_Salary, if (@tmp_num is null
    		OR @tmp_DepartmentId!=e.DepartmentId, @tmp_num:=0, @tmp_num:=@tmp_num) AS tmp_num2, IF(@tmp_DepartmentId=e.DepartmentId
    		AND @tmp_Salary!=0,@tmp_num:=@tmp_num+1,@tmp_num:=@tmp_num) AS d_num, @tmp_DepartmentId:=e.DepartmentId AS tmp_DepartmentId
    FROM 
        (SELECT d.Name,
		 e1.Name AS Employee,
		 e1.Salary,
		 e1.DepartmentId,
		 e1.Id
        FROM Employee e1, Department d
        WHERE e1.DepartmentId = d.Id
        GROUP BY  e1.DepartmentId,Employee
        ORDER BY  e1.DepartmentId asc, e1.Salary desc, e1.Id DESC ) AS e ) AS t
    WHERE d_num < 3
    