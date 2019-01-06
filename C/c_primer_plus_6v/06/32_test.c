#include <stdio.h>

int main(void)
{
	char char_array[255];
	int i = 0;


	printf("Please enter a line: ");
	while (scanf("%c", &char_array[i]) == 1 && char_array[i] != '\n')
	{
		i++;
	}

	for (; i >= 0; i--) 
	{
		printf("%c", char_array[i]);
	}
	printf("\n");

	return 0;
}
