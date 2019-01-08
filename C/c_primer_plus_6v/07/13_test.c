#include <stdio.h>

int main(void)
{
	long c;
	int s, r;
	char ch;
	c = 0L;
	s = r = 0;

	while ((ch = getchar()) != '#')
	{
		switch (ch)
		{
			case '\n': r++; break;
			case ' ': s++; break;
			default: c++; break;
		}
	}

	printf("space = %d, wrap = %d, other char = %ld\n",
			s, r, c);
	printf("Bye!\n");

	return 0;
}
