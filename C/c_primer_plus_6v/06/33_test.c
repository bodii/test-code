#include <stdio.h>
#define FUNDS 1000000
#define RATE 0.08
#define OUT 100000

int main(void)
{
	int year = 1;
	int current_funds = FUNDS;
	do
	{
		printf("The remaining bonus is $%d.\n", current_funds);
		current_funds += current_funds * RATE;
		current_funds -= OUT;
		year++;
	} while (current_funds > 0);
	printf("It takes %d years to withdraw all the funds.\n", year);

	return 0;
}
