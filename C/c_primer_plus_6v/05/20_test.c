/* 用户输入天数，然后将其转换成周数和天数 */
#include <stdio.h>
#define WEEK 7

int main(void)
{

	int days, weeks, rem_days;	
	
	printf("Please enter the number of days: \n");
	scanf("%d", &days);
	while (days > 0)
	{
		weeks = days / WEEK;
		rem_days = days % WEEK;
		printf("%d days are %d weeks, %d days\n", days, weeks, rem_days);
		scanf("%d", &days);
	}
	printf("bye!\n");

	return 0;
}

