#include <stdio.h>
#include <ctype.h>
#include <stdlib.h>
#define NUM 26

void readFile(char * fname);

int alpha(char ch);

int main(void)
{
	char * fname = "abc_test_file";

	readFile(fname);
	
	return 0;
}


void readFile(char * fname)
{
	FILE * fb;
	char ch;


	fb = fopen(fname, "r");

	if (fb == NULL) {
		printf("file error!\n");
		exit(1);
	}

	while ((ch = fgetc(fb)) != EOF)
	{
		printf("%c ", ch);
		if (isalpha(ch))
		{
			printf(" is alpha");
			printf(". in %d", alpha(ch));
		} else {
			printf(" is not alpha");
		}

		printf("\n");
	}
}


int alpha(char ch)
{
	char * alpha = "abcdefghijklmnopqrstuvwxyz"; 
	int i;
	int returnVal = -1;

	for(i = 0; i < NUM; i++) 
	{
		if (tolower(ch) == alpha[i])
		{
			returnVal = i + 1;
			break;
		}
	}
	
	return returnVal;
}
