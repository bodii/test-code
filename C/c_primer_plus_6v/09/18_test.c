#include <stdio.h>

void chline(char ch, int i, int j)
{
	int r, c;

	for (r = 0; r < j; r++)
	{
		for (c = 0; c < i; c++)
			putchar(ch);
		putchar('\n');
	}
}


int main(void)
{
	int i = 6, j = 3;

	chline('*', i, j);

	return 0;
}
