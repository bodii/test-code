#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX 40


int main(void)
{
	FILE *fp;
	char words[MAX];

	if ( (fp = fopen("wordy2", "a+")) == NULL)
	{
		fprintf(stdout, "不能打开\"wordy2\"文件.");
		exit(EXIT_FAILURE);
	}

	puts("请输入一段文字添加到文件中，如果输入的是以#开头，则程序退出");

	while ((fscanf(stdin, "%40s", words)) == 1 && (words[0] != '#'))
		fprintf(fp, "%s\n", words);

	puts("文件的内容是：");
	rewind(fp); /* 返回到文件的开始处 */

	while((fscanf(fp, "%s", words)) == 1)
		puts(words);

	puts("Done!");

	if (fclose(fp) != 0)
		fprintf(stderr, "关闭文件失败!\n");

	return 0;
}
