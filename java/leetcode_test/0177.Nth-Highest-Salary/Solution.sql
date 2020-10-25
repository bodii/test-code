CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
    if N > 0 then
        set N = N - 1;
    else
        RETURN null;
    end if;
  RETURN (
      # Write your MySQL query statement below.
      select Salary from Employee group by Salary order by Salary desc limit N,1
  );
END;