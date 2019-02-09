#include <stdio.h>

#define PR(NAME, NUM) printf("name: %s; value: %d; address: %p\n", NAME, NUM, &NUM)

int main(void)
{
	int v = 23;
	printf("%p\n", &v);
	PR("fop", v);

	return 0;
}
