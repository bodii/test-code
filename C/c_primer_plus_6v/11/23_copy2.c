// strcpy()

#include <stdio.h>
#include <string.h>
#define WORDS "beast"
#define SIZE 40

int main(void)
{
	const char * orig = WORDS;
	char copy[SIZE] = "Be the best that you can be.";
	char * ps;

	puts(orig);
	puts(copy);
	ps = strcpy(copy + 7, orig);
	puts(copy);
	puts(ps);

	return 0;
}

/*
 注意，strcpy()把源字符串中的空字符也拷贝在内。在该例中，空字符覆盖
 了copy数组中that的第1个t。
  orig:"Be the best that you can be.\0"
  copy:"beast\0"

  [strcpy(copy + 7, orig)后]:
  copy:"Be the beast\0hat you can be.\0"
  所以puts(copy);输出 "Be the beast"
*/
