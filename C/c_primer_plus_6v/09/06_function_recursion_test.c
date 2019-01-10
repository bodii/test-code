// 递归演示
#include <stdio.h>

void up_and_down(int);

int main(void)
{
	up_and_down(1);

	return 0;
}

void up_and_down(int n)
{
	printf("Level %d: n location %p\n", n, &n);  // #1
	if (n < 4)
		up_and_down(n + 1);
	printf("LEVEL %d: n location %p\n", n, &n);  // #2
}

/*
当#1执行完成，执行到第4级时，n的值是4,所以if测试条件为假。up_and_down()
不再调用自已。第4级调用接着执行打印语句#2,即打印LEVEL 4, 因为n的值是4.
此时，第4级调用结束，控制被传回它的主调用函数（即第3级调用）,以此类推。
*/
