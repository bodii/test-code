#include <stdio.h>
#define VAR 8


int main(void)
{
	char ch;
	unsigned long s;
	char string[VAR];
	char integer[VAR];
	int i  = 0;

	s = 0L;

	while ((ch = getchar()) != '#')
	{
		if (ch == '\n')
			continue;

		printf("%c=%d; ", ch, ch);

		i++;
		if (i > 0 && (i % VAR) == 0)
			printf("\n");
	}
	printf("Bye!\n");


	return 0;
}
