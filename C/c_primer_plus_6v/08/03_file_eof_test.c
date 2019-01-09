// 输入文件名，读出文件内容
#include <stdio.h>
#include <stdlib.h>

int main(void)
{
	char ch;
	FILE * fb;
	char fname[50];

	printf("Enter the name of the file: ");
	if ((scanf("%s", fname)) == 0)
	{
		printf("Failed to enter. Bye\n");
		exit(1);
	}

	fb = fopen(fname, "r"); // 打开待读取文件
	if (fb == NULL)
	{
		printf("Failed to open file. Bye\n");
		exit(1);
	}

	// 使用getc(fp)从文件中获取一个字符
	while ((ch = getc(fb)) != EOF)
		putchar(ch);

	fclose(fb);   // 关闭文件

	return 0;
}
