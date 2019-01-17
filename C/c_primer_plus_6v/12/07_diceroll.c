#include <stdio.h>
#include <stdlib.h>  /* 包含rand函数的原型 */
#include "diceroll2.h"

int roll_count = 0;

static int rollem(int sides) /* 该函数属于该文件私有 */
{
	int roll; /* 随机数变量 */

	roll = rand() % sides + 1;  /* 获取随机数，并赋给随机数变量 */
	++roll_count; /* 统计随机的次数 */

	return roll;
}


int roll_n_dice(int dice, int sides)
{
	int d;
	int total = 0;

	if (sides < 2)  /* 最少面数不能少于是2 */
	{
		printf("骰子的面数不能少于2!\n");
		return -2;
	}

	if (dice  < 1)
	{
		printf("至少随机一次!\n");
		return -1;
	}

	for (d = 0; d < dice; d++)
		total += rollem(sides);

	return total;
}
