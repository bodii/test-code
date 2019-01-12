// 计算每年的总降水量、年平均降水量和5年中每月的平均降水量
#include <stdio.h>
#define MONTHS 12  // 一年的月份数
#define YEARS 5    // 年数

int main(void)
{
	// 用2010~2014年的降水量数据初始化数组
	const float rain[YEARS][MONTHS] = 
	{
		{ 4.3, 4.3, 4.5, 3.0, 1.3, 5.5, 2.6, 1.7, 3.4, 1.8, 2.2, 3.1 },
		{ 4.3, 4.3, 4.5, 3.0, 1.3, 5.5, 2.6, 1.7, 3.4, 1.8, 2.2, 3.1 },
		{ 4.3, 4.3, 4.5, 3.0, 1.3, 5.5, 2.6, 1.7, 3.4, 1.8, 2.2, 3.1 },
		{ 4.3, 4.3, 4.5, 3.0, 1.3, 5.5, 2.6, 1.7, 3.4, 1.8, 2.2, 3.1 },
		{ 4.3, 4.3, 4.5, 3.0, 1.3, 5.5, 2.6, 1.7, 3.4, 1.8, 2.2, 3.1 },
	};

	int year, month;
	float subtot, total;
	printf(" YEAR  RAINFALL  (inches)\n");
	for (year = 0, total = 0; year < YEARS; year ++)
	{
		// 每一年各月的降水量总和
		for (month = 0, subtot = 0; month < MONTHS; month++)
			subtot += rain[year][month];
		printf("%5d %15.1f\n", 2010 + year, subtot);
		total += subtot; // 5年的总降水量
	}

	printf("\nThe yearly average is %.1f inches.\n\n", total / YEARS);
	printf("MONTHLY ACERAGES:\n\n");
	printf(" Jan Feb Mar Apr May Jun Jul Aug Sep Oct ");
	printf(" Nov Dec\n");

	for (month = 0; month < MONTHS; month++)
	{
		// 每个月, 5年的总降水量
		for (year = 0, subtot = 0; year < YEARS; year++)
			subtot += rain[year][month];
		printf("%4.1f ", subtot / YEARS);
	}
	printf("\n");

	return 0;
}
