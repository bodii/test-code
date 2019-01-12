// 指针操作

#include <stdio.h>

int main(void)
{
	int urn[5] = { 100, 200, 300, 400, 500 };
	int * ptr1, *ptr2, *ptr3;

	ptr1 = urn;   // 把一个地址赋给指针
	ptr2 = &urn[2]; // 把一个地址赋给指针
	// 解引用指针，以及获得指针的地址
	printf("pointer value, dereferenced pointer, pointer address:\n");
	printf("ptr1 = %p, *ptr1 = %d, &ptr1 = %p\n", ptr1, *ptr1, &ptr1);

	// 指针加法
	ptr3 = ptr1 + 4;
	printf("\nvalues after ptr1++:\n");
	printf("ptr1 + 4 = %p, *(ptr1 + 4) = %d\n", ptr1 + 4, *(ptr1 + 4));

	// 递增指针
	ptr1++;
	printf("\nvalues after ptr1++:\n");
	printf("ptr1 = %p, *ptr1 = %d, &ptr1 = %p\n", ptr1, *ptr1, &ptr1);

	// 递减指针
	ptr2--;
	printf("\nvalues after --ptr2:\n");
	printf("ptr2 = %p, *ptr2 = %d, &ptr2 = %p\n", ptr2, *ptr2, &ptr2);

	--ptr1; // 恢复为初始值
	++ptr2; // 恢复为初始值
	printf("\nPointers reset to original values:\n");
	printf("ptr1 = %p, ptr2 = %p\n", ptr1, ptr2);

	// 一个指针减去另一个指针
	printf("\nsubtracting one pointer from another:\n");
	printf("ptr2 = %p, ptr1 = %p, ptr2 - ptr1 = %td\n", ptr2, ptr1, ptr2 - ptr1);

	// 一个指针减去一个整数
	printf("\nsubtracting an int from a pointer:\n");
	printf("ptr3 = %p, ptr3 -2 = %p\n", ptr3, ptr3 - 2);

	return 0;
}
