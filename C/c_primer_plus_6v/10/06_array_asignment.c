// 给数组的元素赋值

#include <stdio.h>
#define SIZE 50

int main(void)
{
	int counter, evens[SIZE];

	for (counter = 0; counter < SIZE; counter++)
		evens[counter] = 2 * counter;
	
	for (counter = 0; counter < SIZE; counter++)
		printf("%d, ", evens[counter]);

	putchar('\n');

	return 0;
}
