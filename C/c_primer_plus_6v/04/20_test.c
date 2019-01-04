// 请输入名和姓（根据英文书写习惯，名和姓中间有一个空格)
#include <stdio.h>

int main(void)
{
	char name[10], surname[10];

	printf("Please enter your name and surname:\n");
	scanf("%s %s", name, surname);
	printf("Your name is: %s %s\n", name, surname);

	return 0;
}
