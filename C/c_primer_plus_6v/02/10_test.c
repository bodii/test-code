#include <stdio.h>

int main(void)
{
	char * name1 = "Gustav";
	char * name2 = "Mahler";

	printf("%s %s\n", name1, name2);
	printf("%s\n%s\n", name1, name2);
	printf("%s ", name1);
	printf("%s\n", name2);

	return 0;
}

