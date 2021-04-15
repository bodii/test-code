with t as(
    select *,id - row_number() over(order by id) as r
    from stadium
    where people >= 100
)

select id,visit_date,people
from t
where r in(
    select r
    from t
    group by r
    having count(r) >= 3
);