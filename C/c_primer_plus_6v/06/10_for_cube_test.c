/* 使用for循环创建一个立方表 */
#include <stdio.h>

int main(void)
{
	int num;

	printf("%5s %5s cubed\n", "n", "n");

	for (num = 1; num <=6; num++)
		printf("%5d %5d\n", num, num * num * num);

	return 0;
}
