// 包含头文件两次

#include <stdio.h>
#include "names2.h"
#include "names2.h" // 不小心第2次包含头文件

int main()
{
	names winner = { "Less", "lsmoor" };
	printf("The winner is %s %s.\n", winner.first, winner.last);

	return 0;
}
