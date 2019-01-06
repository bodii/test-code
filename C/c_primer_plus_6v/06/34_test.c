// 邓巴数 150
#include <stdio.h>
#define INIT 5

int main(void)
{
	int frieds, week;
	frieds = 5;
	week = 0;

	while (frieds < 150) 
	{
		week++;
		frieds -= week;
		frieds *= 2;
		printf("week = %d, frieds = %d\n", week, frieds);
	}

	return 0;
}
