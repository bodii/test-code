select product_id,
min(if(store='store1', price,null)) as store1,
min(if(store='store2', price,null)) as store2,
min(if(store='store3', price,null)) as store3
from Products
group by product_id;
