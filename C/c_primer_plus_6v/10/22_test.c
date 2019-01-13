#include <stdio.h>

void copy_arr(double *, double *, int);

void copy_ptr(double *, double *, int);

void copy_ptrs(double *, double *, double *);

int main(void)
{
	double source[5] = { 1.1, 2.2, 3.3, 4.4, 5.5 };
	double target1[5];
	double target2[5];
	double target3[5];

	copy_arr(target1, source, 5);
	copy_ptr(target2, source, 5);
	copy_ptrs(target3, source, source + 5);


	return 0;
}


/* 使用带数组表示法的函数时行拷贝 */
void copy_arr(double * target1, double * source, int n)
{
	int i;

	for (i = 0; i < n; i++)
		target1[i] = source[i];

	printf("target1: \n");
	for (i = 0; i < n; i++)
		printf("index: %d  value: %.1f\n", i, target1[i]);
}

void copy_ptr(double * target2, double * source, int n)
{
	int i;
	
	for (i=0; i < n; i++)
		target2[i] = *(source + i);

	printf("target2: \n");
	for (i =0; i < n; i++)
		printf("index: %d  value: %.1f\n", i, target2[i]);
}

void copy_ptrs(double * target3, double * source, double * ptrs)
{
	int i = 0;
	int n = 0;

	for (i = 0; i < (ptrs-source); i++)
		target3[i] = *(source + i);

	printf("target3: \n");
	for (i = 0; i < (ptrs-source); i++)
		printf("index: %d  value: %.1f\n", i, target3[i]);

}
