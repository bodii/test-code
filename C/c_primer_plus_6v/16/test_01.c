#include <stdio.h>

#define FEET 4
#define POD1 FEET + FEET
#define POD2 (FEET+FEET)
#define POD3 ((FEET)+(FEET))

int main(int argc, char * argv[])
{
	int plort1, plort2, plort3;

	plort1 = FEET * POD1;
	plort2 = FEET * POD2;
	plort3 = FEET * POD3;
	printf("%d\n", plort1);
	printf("%d\n", plort2);
	printf("%d\n", plort3);

	return 0;
}
