#include <stdio.h>

int main(void)
{
	int guess = 1;
	char response;

	printf("Uh...is your number %d?\n", guess);
	while ((response = getchar()) != 'y') 
	{
		if (response == 'n')
			printf("Well, then, is it %d?\n", ++guess);
		else 
			printf("Sorry, I understand only y or n.\n");
		while (getchar() != '\n')
			continue;
	}
	printf("I knew I could do it!\n");

	return 0;
}
