#include <stdio.h>

/*
 编写一个程序，生成以下输出：
 For he's a jolly good fellow!
 For he's a jolly good fellow!
 For he's a jolly good fellow!
 Which nobody can deny!

 除了main()函数以外，该程序还要调用两个自定义函数：一个名为jolly(), 
 用于打印前3条消息,调用一次打印一条;另一个函数名为deny(), 打印最后一
 条消息。
 */

// 自定义函数 声明
void jolly(void);
void deny(void);

// 自定义函数 实体
void jolly(void)
{
	printf("For he's a jolly good fellow!\n");
}

void deny(void)
{
	printf("Which nobody can deny!\n");
}


int main(void)
{
	jolly();
	jolly();
	jolly();
	deny();

	return 0;
}
