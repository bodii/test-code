#include <stdio.h>

int main(void)
{
	unsigned int un = 3000000000; /* int为32位和short为16位的系统 */
	short end = 200;
	long big = 65537;
	long long verybig = 12345678908642;

	printf("un = %u and not %d\n", un, un);
	printf("end = %hd and %d\n", end, end);
	printf("big = %ld\n", big);
	printf("veruybig = %lld\n", verybig);

	return 0;
}
