#include <stdio.h>

char * mystrncpy(char * , const char *, int);
int mystrlen(const char *);


int main(void)
{
	char * s1 = "abcde";
	char * s2 = "defsdaf";
	int n  = 20;
	char * val;

	val = mystrncpy(s1, s2, n);
	puts(val);

	return 0;
}


char * mystrncpy(char * s1, const char * restrict s2, int n)
{
	int s1_len;
	int i = 0;
	char ptr[n];

	s1_len = mystrlen(s1);
	printf("%d\n", s1_len);

	if(s1_len > 0)
	{
		while (*s1 && n > i)
		{
			ptr[i] = *s1;
			s1++;
			i++;
		}
	}

	if ((mystrlen(s2) + s1_len) < n)
	{
		while ( *s2 )
		{
			ptr[i] = *s2;
			s2++;
			i++;
		}
	}

	s1 = ptr;

	return s1;
}


int mystrlen(const char * str)
{
	int count = 0;

	while (*str)
	{
		str++;
		count++;
	}

	return count;
}
