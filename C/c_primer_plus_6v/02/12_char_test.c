#include <stdio.h>

int main(void)
{
	char * str1 = "this";
	char * str2 = "this is string";
	char str1Arr [] = "this";
	char str2Arr [] = "this is string";

	printf("string1: %s\n", str1);
	printf("string1 sizeof: %ld\n", sizeof(str1));
	printf("string2: %s\n", str2);
	printf("string2 sizeof: %ld\n", sizeof(str2));

	printf("string1 to array: %s\n", str1Arr);
	printf("string1 to array sizeof: %ld\n", sizeof(str1Arr));
	printf("string2 to array: %s\n", str2Arr);
	printf("string2 to array sizeof: %ld\n", sizeof(str2Arr));

	return 0;
}
