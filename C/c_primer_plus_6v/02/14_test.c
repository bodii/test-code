#include <stdio.h>

/*
	编写一个程序，生成以下输出：
	Brazil, Russia, India, China
	India, China,
	Brazil, Russia

	除了main()以外，该程序还要调用两个自定义函数：一个名为br(),
	调用一次打印一次"Brazil, Russia";另一个名为ic(),调用一次打印
	一次"India, China"。其他内容在main()函数中完成。
 */

void br(void)
{
	printf("Brazil, Russia");
}

void ic(void)
{
	printf("India, China");
}

int main(void)
{
	br();
	printf(", ");
	ic();
	printf("\n");
	ic();
	printf(",\n");
	br();
	printf("\n");



	return 0;
}
