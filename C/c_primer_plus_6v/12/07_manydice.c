#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include "diceroll2.h"

int main(void)
{
	int dice, roll;
	int sides;
	int status;

	srand((unsigned int) time(0));  /* 随机数种子 */
	printf("请输入骰子的面数， 0是结束\n");
	while( scanf("%d", &sides) == 1 && sides > 0)
	{
		printf("要随机几次？\n");
		if ( (status = scanf("%d", &dice)) != 1)
		{
			if (status == EOF)
				break;
			else
			{
				printf("请输入一个有效的整数");
				printf("， 并开始想要的随机次数\n");
				while (getchar() != '\n')
					continue;  /* 过滤掉无用的输入字符 */

				printf("请输入骰子的面数，0是结束\n");
				continue;  /* 进入下一次循环 */
			}
		}

		roll = roll_n_dice(dice, sides);
		printf("你得到随机总数是%d 随机了%d次%d面的骰子\n",
				roll, dice, sides);
	}

	printf("使用rollem()函数的次数是 %d\n",
			roll_count);
	printf("祝你好运!\n");
	
	return 0;
}
