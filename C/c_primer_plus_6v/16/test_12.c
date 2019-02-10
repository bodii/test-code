#include <stdio.h>

#define AVER(X, Y) (1 / ((1/X + 1/Y) / 2))

int main(int argc, char * argv[])
{
	char * a1;

	a1 = argv[1];
	printf("%d\n", atoi(a1));

	return 0;
}
