#include <stdio.h>


int main(void) /* 把2音寻（测水深的单位） 转换成英尺 */
{
	int feet, fathoms; // 使用有意义的变量名
	// 使用空行
	fathoms = 2;
	feet = 6 * fathoms; // 每行一条语句

	printf("There are %d feet in %d fathoms! \n", feet, fathoms);

	return 0;
}
