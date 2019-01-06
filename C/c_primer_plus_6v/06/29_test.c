#include <stdio.h>
#include <string.h>

int main(void)
{
	char word[100];
	int i;

	printf("Please enter a word: ");
	for (scanf("%s", word), i=strlen(word); i >= 0; i--)
		printf("%c", word[i]);
	printf("\n");

	return 0;
}

