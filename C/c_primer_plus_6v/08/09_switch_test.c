#include <stdio.h>

char get_choice(void);

void count(void);

int main(void)
{
	char choice;

	while ((choice = get_choice()) != 'q')
	{
		switch (choice)
		{
			case 'a': printf("Buy low, sell high.\n");
					  break;
			case 'b': putchar('\a'); 
					  break;
			case 'c': count();
					  break;
			default: printf("Program error!\n");
					 break;
		}
	}

	return 0;
}


char get_choice(void)
{
	char ch;

	printf("Enter the letter of your choice:\n"
			"%-10s\t\t%-10s\n"
			"%-10s\t\t%-10s\n",
			"a. advice", "b. bell",
			"c. count", "q. quit"
		  );
	ch = getchar();
	while ((ch < 'a' || ch > 'c') && ch !='q')
	{
		printf("Please respond with a, b, c or q.\n");
		ch = getchar();
	}

	return ch;
}


void count(void)
{
	int n, i;
	printf("Count how far? Enter an integer:\n");
	scanf("%d", &n);
	for (i = 1; i <= n; i++)
		printf("%d\n", i);
}

