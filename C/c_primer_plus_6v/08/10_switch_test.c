// 菜单程序
#include <stdio.h>

char get_choice(void);

char get_first(void);

int get_int(void);

void count(void);


int main(void)
{
	int choice;
	void count(void);

	while ((choice = get_choice()) != 'q')
	{
		switch(choice)
		{
			case 'a': printf("Buy6 low, sell high.\n");
					  break;
			case 'b': putchar('\a'); 
					  break;
			case 'c': count();
					  break;
			default: printf("Program error!\n");
					 break;
		}
	}
	printf("Bye.\n");

	return 0;
}


void count(void)
{
	int n, i;

	printf("Count how far? Enter an integer:\n");
	n = get_int();
	for (i = 0; i <= n; i++)
		printf("%d\n", i);
	while (getchar() != '\n')
		continue;
}


char get_choice(void)
{
	int ch;

	printf("Enter the letter of your choice:\n");
	printf("%-10s\t\t%-10s\n%-10s\t\t%-10s\n",
			"a. advice", "b. bell", 
			"c. count", "q. quit");
	ch = get_first();
	while ((ch < 'a' || ch > 'c') && ch != 'q')
	{
		printf("Please respond whit a, b, c, or q.\n");
		ch = get_first();
	}
	return ch;
}


char get_first(void)
{
	char ch;

	ch = getchar(); // 获取第一个
	while ( getchar() != '\n') // 过滤掉剩余
		continue;

	return ch;
}

int get_int(void)
{
	int input;
	char ch;

	while ((scanf("%d", &input)) != 1)
	{
		while ((ch = getchar()) != '\n')
			putchar(ch); // 处理错误输出
		printf(" is not an integer.\nPlease enter an ");
		printf("integer value, such as 25, -178, or 3: ");
	}
	return input;
}


