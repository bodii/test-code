# Write your MySQL query statement below

select Score, cast(Ranks as signed) as `Rank` from
(select s.Score,
@r:=(case when @p=s.Score && @r >= 1 then @r else @r+1 end) as Ranks,
@p:=s.Score as tmp_score 
from
(select Score, Id from Scores  order by Score desc) as s,
(select @r:=0, @p:=0) as t) as v;