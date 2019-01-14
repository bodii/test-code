// 使用缩短字符串长度的函数

#include <stdio.h>
#include <string.h>  /* 内含字符串函数原型 */

void fit(char *, unsigned int);

int main(void)
{
	char mesg [] = "Things should be as simple as possible,"
					" but not simpler.";

	puts(mesg);
	fit(mesg, 38);
	puts(mesg);
	puts("Let's look at some more of the string.");
	puts(mesg + 39);
	
	return 0;
}

void fit(char * string, unsigned int size)
{
	if (strlen(string) > size)
		string[size] = '\0';
}


/*
 fit()函数的把第39个元素的逗号替换成'\0'字符。puts()函数在空字符停止
 输出，并忽略其余字符。然而，这些字符还在缓冲区中，下面的函数调用把这
 些字符打印了出来：
 puts(mesg + 8);
 表达式mesg + 39是mesg[39]的地址，该地址上储存的是空格字符。所以put()
 显示该字符并继续输出直至遇到原来字符串中的空字符。
 原始字符串：
 Hold on to your hats, hackers.\0
 调用fit(mesg, 7)之后的字符串
 Hold on\0to your hasts, hackers.\0
*/
