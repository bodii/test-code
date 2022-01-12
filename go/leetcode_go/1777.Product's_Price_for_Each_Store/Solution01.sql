# Write your MySQL query statement below

select product_id,
min(case store when 'store1' then price else null end) as store1,
min(case store when 'store2' then price else null end) as store2,
min(case store when 'store3' then price else null end) as store3
from Products
group by product_id;
