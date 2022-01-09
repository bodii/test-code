# Write your MySQL query statement below

select e1.employee_id, count(1) as team_size from 
Employee as e1 
left join Employee as e2
on e1.team_id=e2.team_id
group by e1.employee_id;
