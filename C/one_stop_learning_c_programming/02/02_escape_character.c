#include <stdio.h>

/* 转义字符节 */

int main(void)
{
	/* linux: 换行 \n */
	/* windows: 换行 \r\n */
	printf("Single quote: \\' = \' \n");
	printf("double quote: \\\" = \" \n");
	printf("question quote: \\? =  \? \n");
	printf("backslash quote: \\\\ = \\ \n");
	printf("alter quote: \\a =  \a \n");
	printf("backspace: \\b = \b \n");
	printf("form feed: \\f = \f \n");
	printf("line feed: \\n = \n \n");
	printf("carriage return: \\r = \r \n");
	printf("horizontal tab: \\t = \t \n");
	printf("vertical tab: \\v = \v \n");
}

