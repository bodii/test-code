#include <stdio.h>

/*
 编写一个程序，生成以下格式的输出：
 Smile!Smile!Smile!
 Smile!Smile!
 Smile!

 该程序要定义一个函数，该函数被调用一次打印一次"Smile!",根据程序的需要使用
 该函数
 */

void smile(void)
{
	printf("Smile!");
}

int main(void)
{
	smile();smile();smile();printf("\n");
	smile();smile();printf("\n");
	smile();printf("\n");

	return 0;
}
