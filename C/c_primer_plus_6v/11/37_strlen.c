// 编写strlen函数

#include <stdio.h>

int strlen2(char *); // 获取字符串长度

int main(void)
{
	char str[] = "abcdf dfds w2r !flfd j dfsf";

	printf("\"%s\"  length is: %d\n", str, strlen2(str)); 
	return 0;
}


int strlen2(char * st)
{
	int count = 0;

	while ( st[count] != '\n' && st[count] != '\0')
		count++;

	return count;
}
