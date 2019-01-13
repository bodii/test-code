// 使用fgets()和fputs()

#include <stdio.h>
#define STLEN 10

int main(void)
{
	char words[STLEN];

	puts("Enter strings (empty line to quit):");
	while (fgets(words, STLEN, stdin) != NULL && words[0] != '\n')
		fputs(words, stdout);
	puts("Done.");

	return 0;
}

/*
 程序中的fgets()一次读入STLEN -1 个字符（该例中为9个字符).所以，一开始
 它只读入了"By the wa", 并储存为By the wa\0;接着fputs()打印该字符串，而
 且并未换行。然后while循环进入下一轮迭代，fgets()继续从剩余的输入中读入
 数据，即读入"y, the ge" 并储存为y, the ge\0;接着fputs()在刚才打印字符串
 的这一行接着打印第2次读入的字符串。然后while进入下一轮迭代，fgets()继续
 读取输入、fputs()打印字符串，这一过程循环进行，直到读入最后的"tion\n".
 fgets()将其储存为tion\n\0,fputs()打印该字符串，由于字符串中的\n,光标被
 移至下一行开始处。
 */
