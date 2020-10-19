# Write your MySQL query statement below
/*
# 1
SELECT 
    MAX(Salary) AS SecondHighestSalary
FROM
    Employee
WHERE
    Salary < (SELECT 
            MAX(Salary) AS ms
        FROM
            Employee
        LIMIT 1)
LIMIT 1;
*/


/*
# 2
SELECT 
    IFNULL((SELECT 
                    Salary
                FROM
                    Employee
                WHERE
                    Salary NOT IN (SELECT 
                            MAX(Salary) AS ms
                        FROM
                            Employee)
                ORDER BY Salary DESC
                LIMIT 1),
            NULL) AS SecondHighestSalary;
*/

# 3
SELECT 
    IFNULL((SELECT 
                    Salary
                FROM
                    Employee
                GROUP BY Salary
                ORDER BY Salary DESC
                LIMIT 1 , 1),
            NULL) AS SecondHighestSalary;


# 4
SELECT 
    IFNULL((SELECT DISTINCT
                    Salary
                FROM
                    Employee
                ORDER BY Salary DESC
                LIMIT 1 OFFSET 1),
            NULL) AS SecondHighestSalary;

