# Write your MySQL query statement below

select e1.employee_id, e2.team_size from Employee as e1 
left join (select team_id, count(team_id) as team_size from Employee group by team_id) e2 
on e1.team_id=e2.team_id;
