
#include <stdio.h>

char * strEmpty(char *);

int main(void)
{
	char * str = " sdfdsdsaf dfdfaafdf DFFD";

	if (strEmpty(str))
		puts(str);
	else 
		printf("The first character of \"%s\" is empty.\n",
				str);

	return 0;
}


char * strEmpty(char * st)
{
	char * ret_val;
	int i = 0;

	if (st[i] != '\0' && st[i] != '\n' && st[i] != ' ')
		ret_val = st;
	return ret_val;
}
