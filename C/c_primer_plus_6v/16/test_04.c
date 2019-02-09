/* 定义EVEN_GT(X, Y)宏，如果X为偶数且大于Y, 该宏返回1. */

#include <stdio.h>

#define EVEN_GT(X, Y) ((X % 2 == 0) && (X > Y)?1:0)
#define SEC_MSG "是偶数并且大于"
#define ERR_MSG "非偶数或不大于"


int main(int argc, char * argv[])
{
	//PR(4, 7);
	char v[24];
	sprintf(v, "%d%s%d", 4, EVEN_GT(4, 6) ? SEC_MSG : ERR_MSG, 6);
	puts(v);

	sprintf(v, "%d%s%d", 8, EVEN_GT(8, 3) ? SEC_MSG : ERR_MSG, 3);
	puts(v);

	return 0;
}
