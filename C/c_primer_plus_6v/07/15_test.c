#include <stdio.h>
#include <stdbool.h>

int main(void)
{
	int is;
	long count = 0L;
	long scount = 0L;
	bool stop = false; 

	while(scanf("%d", &is) == 1 && !stop)
	{
		if (is == 0)
		{
			stop = true;
			break;
		}

		if ( (is % 2) == 0 )
		{
			count++;
			printf("%ld", count);
			printf("\t%d\n", is * is);
		} 
		else if ( (is % 2) != 0 ) 
		{
			scount++;
			printf("%ld", scount);
			printf("\t%d\n", is * is);
		}
	}

	return 0;
}

