/*
 假设data1是内含100个double类型元素的数组,data2是内含300个
double类型元素的数组。
a.编写memcpy()的函数调用,把data2中的前100个元素拷贝到data1中。
b.编写memcpy()的函数调用,把data2中的后100个元素拷贝到data1中。
*/

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void fallarray(double ar[], int n);
void showarray(const double ar[], int n);

int main(void)
{
	int data1_num = 100, data2_num = 300;
	double data1[data1_num];
	double data2[data2_num];

	puts("data1 dobule type of array:");
	fallarray(data1, data1_num);
	showarray(data1, data1_num);

	puts("\ndata2 dobule type of array:");
	fallarray(data2, data2_num);
	showarray(data2, data2_num);

	puts("\n把data2中的前100个元素拷贝到data1中");
	memcpy(data1, data2, 100 * sizeof(double));
	showarray(data1, data1_num);

	return 0;
}

void fallarray(double ar[], int n)
{
	int index;

	for (index = 0; index < n; index++)
		ar[index] = (double)rand() / ((double)rand() + 0.1);
}


void showarray(const double ar[], int n)
{
	int index;

	for (index = 0; index < n; index++)
	{
		printf("%9.4f ", ar[index]);
		if (index % 8 == 0)
			putchar('\n');
	}

	if ( index % 8 != 0 )
		putchar('\n');
}


