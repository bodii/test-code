select customer_id, COUNT(1) count_no_trans
	from Visits
 where visit_id not in (select visit_id as vid from Transactions)
 group by customer_id;

