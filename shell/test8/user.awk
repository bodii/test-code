# 在主循环开始以前打印标题，使下面的数据的含义更加明确
# echo 'title to clarify the following data bafore main loop begin'
BEGIN {
	printf "%-19s%s\n","USERNAME","UID";
	printf "-----------------------\n";
}

# 主循环开始运行
# 从/etc/passwd 文件中提取出用户名和UID
{
	printf "%-19s%s\n",$1,$3;
}
