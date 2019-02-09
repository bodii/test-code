#include <stdio.h>

#define NEW1(X) X + 5
#define NEW2(X) X+5
#define NEW3(X) (X)+5
#define NEW4(X) ((X)+5)

int main(int argc, char * argv[])
{
	int y = 3;
	y = NEW1(y);
	printf("NEW1: %d\n", y);

	y = 3;
	y = NEW2(y);
	printf("NEW2: %d\n", y);

	y = 3;
	y = NEW3(y);
	printf("NEW3: %d\n", y);

	y = 3;
	y = NEW4(y);
	printf("NEW4: %d\n", y);

	return 0;
}
