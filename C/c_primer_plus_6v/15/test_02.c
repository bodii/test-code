#include <stdio.h>
#include <string.h>

void to_bin(unsigned int,  unsigned int);
int strToBin(char * str);


int main(int argc, char * argv[])
{
	printf("%d", strToBin(argv[1]));

	return 0;
}


void to_bin(unsigned int bin_str1, unsigned int bin_str2)
{
	printf("%d\n", ~bin_str1);
	printf("%ls\n", &bin_str2);
}

int strToBin(char * str)
{
	int val = 1;
	int i  = 0;

	printf("strlen: %ld\n", strlen(str));
	while(i < strlen(str))
	{
		if (str[i] == '1')
			val |= val<<1;
		else
			val |= val>>1;
		printf("%d\n", val);
		i++;
	}

	return val;
}
