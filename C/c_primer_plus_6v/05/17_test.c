#include <stdio.h>
#define FORMAT "%s! C is coll!\n"

int main(void)
{
	int num = 10;

	printf(FORMAT, FORMAT);
	printf("%d\n", ++num);
	printf("%d\n", num++);
	printf("%d\n", num--);
	printf("%d\n", num);

	return 0;
}
