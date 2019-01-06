#include <stdio.h>

int main(void)
{
	int l, u, i, input_status;
	long s;
    printf("Enter lower and upper integer limits: "); 
	input_status = scanf("%d %d", &l, &u);
	input_status = (input_status == 2) && (l < u);
	while ( input_status ) {
		s = 0L;
		for (i = l; i <= u; i++)
		{
			s += i * i;
		}
		printf("The sums of the squares from %d to %d is %ld\n",
				l * l, u * u, s);
		printf("Enter lower and upper integer limits: "); 
		input_status = scanf("%d %d", &l, &u);
		input_status = (input_status == 2) && (l < u);
	}
	printf("Done.\n");
	return 0;
}
