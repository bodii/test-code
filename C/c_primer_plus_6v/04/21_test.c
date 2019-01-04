// 测试
#include <stdio.h>
#include <string.h>
#define Q "His hamlet was funny without being vulgar."

int main(void)
{
	printf("He sold the painting for $%2.2f.\n", 2.345e2);

	printf("%c%c%c\n", 'H', 105, '\41');

	printf("%s\nhas %ld characters.\n", Q, strlen(Q));

	printf("Is %2.2e the same as %2.2f?\n", 1201.0, 1201.0);

	return 0;
}
