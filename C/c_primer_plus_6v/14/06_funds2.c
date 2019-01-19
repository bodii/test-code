// 传递指向结构的指针

#include <stdio.h>
#define FUNDLEN 50

struct funds {
	char bank[FUNDLEN];
	double bankfund;
	char save[FUNDLEN];
	double savefund;
};

double sum(const struct funds *);  /* 参数是一个指针 */


int main(void)
{
	struct funds stan = {
		"Garlic-Melon Bank",
		4032.27,
		"Lucky's Savings and Loan",
		8543.94
	};

	printf("Stan has a total $%.2f.\n", sum(&stan));

	return 0;
}


double sum(const struct funds * money)
{
	return(money->bankfund + money->savefund);
}
