// 使用 fprintf()、fscanf()和rewind()

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX 41


int main(void)
{
	FILE *fp;
	char words[MAX];

	if ( (fp = fopen("wordy", "a+") ) == NULL )
	{
		fprintf(stdout, "Can't open \"wordy\" file.\n");
		exit(EXIT_FAILURE);
	}

	puts("Enter words to add to the file; press the #");
	puts("key at the beginning of a line to terminate.");

	while ( (fscanf(stdin, "%40s", words) == 1) && (words[0] != '#'))
		fprintf(fp, "%s\n", words); /* fprintf支持标准输入/输出和文件的输入/输出 */
		/* 这里是把输入的内容按每个单词逐行打印到fp指向的文件中 */
	puts("File contents:");
	rewind(fp);  /* 返回到文件开始处 */

	while ( fscanf(fp, "%s", words) == 1) 
		puts(words);

	puts("Done!");

	if (fclose(fp) != 0)
		fprintf(stderr, "Error close file\n");

	return 0;
}
