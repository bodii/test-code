#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>

int main(void)
{
	char ch;
	char * fname = "essayct";
	FILE * fb;
	int count;
	count = 0;

	fb = fopen(fname, "a+");
	if (fb == NULL)
		exit(1);

	while ((ch = getc(fb)) != EOF)
	{
		if(!isspace(ch))
		{
			putchar(ch);
			count++;
		}
	}

	fwrite(&count, sizeof(int), 10, fb);
	fclose(fb);

	return 0;
}
