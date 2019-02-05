/* 编写一个函数，把二进制字符串转换为一个数值 */
#include <stdio.h>
#include <string.h>

int charToInt(char * );
int pow2(int, int);
void to_bin(int);

int main(void)
{
	char * pbin = "01001001";
	int val = 0;

	val = charToInt(pbin);
	printf("%d\n", val);
	to_bin(25);

	return 0;
}


int charToInt(char * ch)
{
	int i = 0, ret_val = 0;

	while(strlen(ch) > i)
	{
		ret_val *= 2;
		if (ch[i] == '1')
			ret_val += 1;
		i++;
	}

	return ret_val;
}

void to_bin(int x)
{
	printf("%d", x % 2);
	if (x > 2)
		to_bin(x / 2);
	else
		putchar('\n');
}
