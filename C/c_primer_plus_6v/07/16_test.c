#include <stdio.h>
#define BASIC_WAGE 1000    // 基本工资 1000美元/小时
#define OVERTIME_40_h 1.5  // 加班超过40小时 = 1.5倍的时间
#define T 0.15  // 税率： 前300美元为15%
#define O 0.20   // 税率： 续150美元为20%
#define S 0.25   // 税率： 余下为25%


int main(void)
{
	float h;
	double wage, tax;
	wage = tax = 0.0;

	if ((scanf("%f", &h)) == 1) 
	{
		if ( h > 40 )
		{
			h *= OVERTIME_40_h;
		}

		wage = h * BASIC_WAGE;
		
		if ( wage <= 300 )
		{
			tax = wage * T;
		}
		else if ( wage > 300 && wage <= 450)
		{
			tax = 300 * T + (wage - 300) * O;
		}
		else if (wage > 450)
		{
			tax = 300 * T + 150 * O + (wage - 450) * S;
		}

		printf("wage = %g; tax = %g; income = %g\n", wage, tax, wage - tax);

	}
	printf("End.\n");
	return 0;
}
