#### 数组和链表

虽然结构不能含有与本身类型相同的结构，但是可以含有指向同类型结构的指针。这种定义是定义链表
(linked list)的基础，链表中的每一项都包含在何处能找到下一项的信息。

假设用户输入的片名是Modern Times，等级为10。程序将为film类型的结构分配空间，把字符串Modern
 Times拷贝到结构中的title成员中，然后设置rating成员为10。为了表明该结构后面没有其他结构，程
 序要把next成员指针设置为NULL(NULL是一个定义在stdio.h头文件中的符号常量，表示空指针)。当然，
 还需要一个单独的指针储存第1个结构的地址，该指针被称为头指针(head pointer).头指针指向链表中
 的第1项。
 ```c
#define TSIZE 45
struct film {
	char title[TSIZE];
	int rating;
	struct film * next;
};

struct film * head;
```
现在，假设用户输入第2部电影及其评级，如Midnight in Paris和8。程序为第2个film类型结构分配空间，
把新结构的地址储存在第1个结构的next成员中(擦写了之前储存在该成员的NULL)，这样链表中第1个结构
中的next指针指向第2个结构。然后程序把Midnight in Paris和8拷贝到新结构中，并把第2个结构中的next
成员设置为NULL，表明该结构是链表中的最后一个结构。

每加入一部新电影，就以相同的方式来处理。新结构的地址将储存在上一个结构上，新信息储存在新结构中，
而且新结构中的next成员设置为NULL。

假设要显示这个链表， 每显示一项，就可以根据该项中已储存的地址来定位下一个待显示的项。然而，这种
方案能正常运行，还需要一个指针储存链表中第1项的地址，因为链表中没有其他项储存该项的地址。此时，
头指针就派上了用场。



#### 使用链表

1. 显示链表
显示链表从设置一个指向第1个结构的指针(名为current)开始。由于头指针(名为head)已经指向链表中的第1
个结构，所以可以用下面的代码来完成:
```c
current = head;
```
然后，可以使用指针表示法访问结构成员：
```c
printf("Movie: %s Rating: %d\n", current->title, current->rating);
```
下一步是根据储存在该结构中next成员中的信息，重新设置current指针指向链表中的下一个结构:
```c
current = current->next;
```
完成这些之后，再重复整个过程。当显示到链表中最后一个项时，current将被设置为NULL,因为这是链表最后
一个结构中next成员的值。
```c
while (current != NULL)
{
	printf("Movie: %s Rating: %d\n", current->title, current->rating);
	current = current->next;
}
```
遍历链表时，为何不直接使用head指针，而要重新创建一个新指针(current)?因为如果使用head会改变head中
的值，程序就找不到链表的开始处。


2. 创建链表
创建链表涉及下面3步：
(1) 使用malloc()为结构分配足够的空间;
(2) 储存结构的地址;
(3) 把当前信息拷贝到结构中。
如无必要不用创建一个结构，所以程序使用临时存储区(input数组)获取用户输入的电影名。如果用户通过键盘
模拟EOF或输入一行空行，将退出下面的循环:
```c
while(s_gets(input, TSIZE) != NULL && input[0] != '\0')
```
如果用户进行输入，程序就分配一个结构的空间，并将其地址赋给指针变量current:
```c
current = (struct film *) malloc(sizeof(struct film));
```
链表中第1个结构的地址应储存在指针变量head中。随后每个结构的地址应储存在其前一个结构的next成员中。因
此，程序要知道它处理的是否是第1个结构。最简单的方法是在程序开始时，把head指针初始化为NULL.然后，程序
可以使用head的值进行判断:
```c
if (head == NULL)   /* 第1个结构 */
	head = current;
else
	prev->next = current;
```
在上面的代码中，指针prev指向上一次分配的结构。

接下来，必须为结构成员设置合适的值。尤其是，把next成员设置为NULL，表明当前结构是链表的最后一个结构。还
要把input数组中的电影名拷贝到title成员中，而且要给rating成员提供一个值：
```c
current->next = NULL;
strcpy(current->title, input);
puts("Enter your rating <0-10>:");
scanf('%d", &current->rating);
```
由于s_gets()限制了只能输入TSIZE-1个字符，所以用strcpy()函数把input数组中的字符串拷贝到title成员很安全。

最后，要为下一次输入做好准备。尤其是，要设置prev指向当前结构。因为在用户输入下一部电影且程序为新结构
分配空间后，当前结构将成为新结构的上一个结构，所以程序在循环末尾这样设置该指针:
```c
prev = current;
```


3. 释放链表
在许多环境中，程序结束时都有会自动释放malloc()分配的内存。但是，最好还是成对调用malloc()和free().因此，
程序在清理内存时为每个已分配的结构都调用了free()函数:
```c
current = head;
while (current != NULL)
{
	current = head;
	head = current->next;
	free(current);
}
```
