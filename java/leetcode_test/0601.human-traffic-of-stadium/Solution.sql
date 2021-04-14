SELECT s2.id,
		 s2.visit_date,
		 s2.people
FROM 
    (SELECT id,
		top_num2,
		max(tmp_index) AS max_index
    FROM 
        (SELECT id,
		 people,
		 if (@tmp_index is null
        		OR people<100,@tmp_index:=0, @tmp_index:=@tmp_index) AS tmp1, if (people>=100, @tmp_index:=@tmp_index+1, @tmp_index:=@tmp_index) AS tmp_index, if (@top_num is null, @top_num:=0, @top_num:=@top_num) AS top_num, if (@tmp_index=0, @top_num:=@top_num+1, @top_num:=@top_num) AS top_num2
        FROM Stadium ) AS s
        WHERE s.tmp_index > 0
        GROUP BY  top_num2
        HAVING max_index>=3 ) AS s1, Stadium AS s2
    WHERE s2.id
	BETWEEN s1.id
		AND s1.id+s1.max_index-1 ;