select product_id,
sum(if(store='store1', price,null)) as store1,
sum(if(store='store2', price,null)) as store2,
sum(if(store='store3', price,null)) as store3
from Products
group by product_id;
