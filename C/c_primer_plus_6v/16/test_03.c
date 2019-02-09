/* 定义一个宏函数，返回两值中的较小值 */
#include <stdio.h>

#define MIN(X, Y) X < Y ? X : Y

int main(int argc, char * argv[])
{
	printf("7 and 3 who are the min: %d\n", 
			MIN(7, 3));
	printf("9 and 12 who are the min: %d\n",
			MIN(9, 12));

	return 0;
}
