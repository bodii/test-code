/* 使用qsort()对其1000个int类型元素的数组，按降序排序,比较函数comp() */

#include <stdio.h>
#include <stdlib.h>

#define LIM 1000

void intArray(int ar[], int n);
void showArray(const int ar[], int n);
int comp(const void * p1, const void * p2);

int main(void)
{
	int ar[LIM];

	intArray(ar, LIM);
	puts("Random list:");
	showArray(ar, LIM);
	qsort(ar, LIM, sizeof(int), comp);
	puts("\nSorted list:");
	showArray(ar, LIM);

	return 0;
}

void intArray(int ar[], int n)
{
	int index;

	for (index = 0; index < n; index++)
		ar[index] = (int)rand();
}

void showArray(const int ar[], int n)
{
	int index;

	for (index = 0; index < n; index++)
	{
		printf("%9d ", ar[index]);

		if (index % 10 == 9)
			putchar('\n');
	}

	if (index % 6 != 0)
		putchar('\n');
}

/* 按从大到小排序 */
int comp(const void * p1, const void * p2)
{
	const int * a1 = (const int *) p1;
	const int * a2 = (const int *) p2;

	if (*a1 > *a2)
		return -1;
	else if (*a1 == *a2)
		return 0;
	else
		return 1;
}
