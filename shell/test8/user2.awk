#printf "title to clarify the following data before main loop begin"
BEGIN{
	printf "%-19s%s\n","USERNAME","UID";
	printf "-----------------------\n";
}

END{
	print "数据已经完成了！"
}
# 主循环开始运行
# 从/etc/passwd 文件中提取出用户名和UID
{
	printf "%-19s%s\n",$1,$3;
}

