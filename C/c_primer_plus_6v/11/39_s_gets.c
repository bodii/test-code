#include <stdio.h>

char * s_gets(char *, int);

int main(void)
{
	char * str; 
	char * s1;
	int limit = 20;

	puts("Please enter a string:");
	s1 = s_gets(str, limit);

	if (s1)
		puts(s1);
	else
		puts("empty.");

	return 0;
}


char * s_gets(char * st, int n)
{
	char * ret_val;
	int i;

	ret_val = fgets(st, n, stdin);

	if (ret_val)
	{
		while (*st != '\n' && *st != '\0')
			st++;
		if (*st == '\n')
			*st = '\0';
		else 
			while (getchar() != '\n')
				continue;
	}

	return ret_val;
}
