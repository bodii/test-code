select
t.Request_at as 'Day',
round(count(if(t.Status!='completed',1,NULL))/count(1),2) as 'Cancellation Rate'
from Trips at t 
inner join Users as u1
on t.Client_Id = u1.Users_Id
and u1.Banned = 'No'
and t.Request_at between '2013-10-01' and '2013-10-03'
inner join Users as u2
on t.Driver_Id = u2.Users_Id
and u2.Banned = 'No'
group by t.Request_at order by Day asc;