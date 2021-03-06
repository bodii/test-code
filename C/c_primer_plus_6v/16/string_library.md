#### string.h库中的memcpy()和memmove()
不能把一个数组赋给另一个数组，所以要通过循环把数组中的每个元素赋给另一个数组相应的元素。
有一个例外的情况是：使用strcpy()和strmcpy()函数来处理字符数组。
memcpy()和memmove()函数提供类似的方法处理任意类型的数组。
下面是这个两个函数的原型：
```c
void *memcpy(void * restrict s1, const void * restrict s2, size_t n);
void * memmove(void *s1, const void *s2, size_t n);
```
这两个函数都从s2指向的位置拷贝n字节到s1指向的位置，而且都返回s1的值。
所不同的是，memcpy()的参数带关键字restrict, 即memcpy()假设两个内存区域之间没有重叠;
而memmove()不作这样的假设，所以拷贝过程类似于先把字节拷贝到一个临时缓冲区，然后再拷贝
到最终目的地。
如果使用memcpy()时，两区域出现重叠会怎样？其行为是未定义的，这意味着该函数可能正常工作，
也可能失败。

由于这两个函数设计用于处理任何数据类型，所有它们的参数都是两个指向void的指针。
C允许把任何类型的指针赋给void *类型的指针。如此宽容导致函数无法知道待拷贝数据的类型。因此，
这两个函数使用第3个参数指明待拷贝的字节数。
注意，对数组而言，字节数一般与元素个数不同。如果要拷贝数组中10个double类型的元素，要使用
10 * sizeof(double),而不是10.


