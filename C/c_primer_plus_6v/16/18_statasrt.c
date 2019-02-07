/* _Static_assert()的使用 */

#include <stdio.h>
#include <limits.h>
#include <assert.h>

/* 这里编译时会报错 */
// _Static_assert(CHAR_BIT == 16, "16-bit char falsely assumed");

int main(void)
{
	puts("char is 16 bits.");
	/* 这里会在运行时报错 */
	assert(CHAR_BIT == 16);

	return 0;
}
