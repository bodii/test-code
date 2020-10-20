# Write your MySQL query statement below

# 1
select c.Name as Customers
from Customers as c left join Orders as o on c.Id=o.CustomerId
where o.CustomerId is null;

# 2 
select Name as Customers
from Customers
where id not in (select CustomerId
from Orders);