#include <stdio.h>

int main(void)
{
	char name[40] = "Thomson";
	float cash;

	cash = 2.25;

	printf("The %s family just may be $%.2f dollars richer!\n", name, cash);

	return 0;
}
