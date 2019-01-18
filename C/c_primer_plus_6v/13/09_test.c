#include <stdio.h>
#include <stdlib.h>

int main(void)
{
	FILE *fp1, *fp2;
	char ch;

	
	if ((fp1 = fopen("terky", "r")) == NULL)
	{
		fprintf(stderr, "open %s die.", "terky");
		exit(EXIT_FAILURE);
	}

	if ((fp2 = fopen("jerky", "w")) == NULL)
	{
		fprintf(stderr, "open %s die.", "jerky");
		exit(EXIT_FAILURE);
	}

	while ((ch = getc(fp1)) != EOF)
	{
		fprintf(stdout, "%c\n", ch);
		putc(ch, fp2);
	}

	fclose(fp1);
	fclose(fp2);

	return 0;
}
