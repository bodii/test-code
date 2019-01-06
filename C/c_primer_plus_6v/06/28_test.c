#include <stdio.h>

int main(void)
{
	char n = 'A';
	char i, s, string;

    scanf("%c", &string);

	for (i = n;  i <= string; i++)
	{
		for (s = i; s < string; s++)
		{
			printf(" ");
		}
		for (s = n; s < i; s++)
		{
			printf("%c",s); 
		}
		for (s = i; s >= n; s--)
		{
			printf("%c", s);
		}
		printf("\n");
	}

	return 0;
}
