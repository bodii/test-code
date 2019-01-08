#include <stdio.h>
#define R1 8.75
#define R2 9.33
#define R3 10.00
#define R4 11.20

void printTop(void); // 打印头部方法

int main(void)
{
	int s;
	float rate;

	printTop();
	while ( scanf("%d", &s) == 1 ) 
	{
		if (s < 1 || s > 5)
			continue;
		
		if ( s == 5 )
			break;

		switch (s) 
		{
			case 1: 
				rate = R1; break;
			case 2:
				rate = R2; break;
			case 3:
				rate = R3; break;
			case 4:
				rate = R4; break;
		}

		printf("Your desired pay rate: %.2f\n", rate);

		printTop();
	}

	printf("end.\n");

	return 0;
}

void printTop(void)
{
	printf("Enter the number corresponding to the "
			"desired pay rate or action:\n");
	printf("1) $%.2f/hr\t\t\t2) $%.2f/hr\n"
			"3) $%.2f/hr\t\t\t4) $%.2f/hr\n"
			"5) quit\n", 
			R1, R2, R3, R4
			);
}

