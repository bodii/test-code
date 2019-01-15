#### 其他字符串函数
ANSI C库有20多个用于处理字符串的函数：

* char * strcpy(char * restrict s1, const char * restrict s2);
该函数把s2指向的字符串（包括空字符)拷贝至s1指向的位置，返回值是s1.

* char * strncpy(char * restrict s1, const char * restrict s2, size_t n);
该函数把s2指向的字符串拷贝至s1指向的位置，拷贝的字符数不超过n,其返回值是s1.
该函数不会拷贝空字符后面的字符，如果源字符串的字符少于n个，目标字符串就以拷
贝的空字符结尾;如果源字符串有n个或超过n个字符，就不拷贝空字符。

* char * strcat(char * restrict s1, const char * restrict s2);
该函数把s2指向的字符串拷贝至s1指向的字符串末尾。s2字符串的第1个字符将覆盖s1
字符串末尾的空字符。该函数返回s1.

* char * strncat(char * restrict s1, const char * restrict s2, size_t n);
该函数把s2字符串中的n个字符拷贝至s1字符串末尾。s2字符串的第1个字符将覆盖
s1字符串末尾的空字符。不会拷贝s2字符串中空字符和其后的字符，并在拷贝字符的
末尾添加一个空字符。该函数返回s1.

* int strcmp(const char * s1, const char * s2);
如果s1字符串在机器排序序列中位于s2字符串的后面，该函数返回一个正数;如果两个
字符串相等，则返回0;如果s1字符串在机器排序序列中位于s2字符串的前面，则返回一
个负数。

* int strncmp(const char * s1, const char * s2, size_t n);
该函数的作用和strcmp()类似，不同的是，该函数在比较n个字符后或遇到第1个空字符
时停止比较。

* char * strchar(const char * s, int c);
如果s字符串中包含c字符，该函数返回指向s字符串首位置的指针(末尾的空字符也是字
符串的一部分，所以在查找范围内);如果在字符串s中未找到c字符，该函数则返回空指
针。

* char * strpbrk(const char * s1, const char * s2);
如果s1字符中包含s2字符串中的任意字符，该函数返回指向s1字符串首位置的指针;如果
在s1字符串中未找到任何s2字符串中的字符，则返回空字符。

* char * strrchr(const char * s, int c);
该函数返回s字符串中c字符的最后一次出现的位置(末尾的空字符也是字符串的一部分，所
以在查找范围内).如果未找到c字符，则返回空指针。

* char * strstr(const char *s1, const char * s2);
该函数返回指向s1字符串中s2字符串出现的首位置。如果在s1中没有找到s2,则返回空指针。

* size_t strlen(const char * s);
该函数返回s字符串中的字符数，不包括末尾的空字符。

restrict 该关键字限制了函数参数的用法。例如，不能把字符串拷贝给本身。

size_t类型是sizeof运算符返回的类型。
C规定sizeof运算符返回一个整数类型，但是并未指定是哪种整数类型，所以size_t在一个系统
中可以是unsigned int, 而在另一个系统中可以是unsigned long。string.h头文件针对特定系
统定义了size_t,或者参考其他有size_t定义的头文件。
我们自定义的s_gets()函数通过while循环检测换行符。其它，这里可以用strchr()代替s_gets()
首先，使用strchr()查找换行符(如果有的话).如果该函数发现了换行符，将返回该换行符的地
址，然后便可用空字符替换该位置上的换行符:
```c
char line[80];
char * find;

fgets(line, 80, stdin);
find = strchar(line, '\n');  // 查找换行符
if (find)
	*find = '\0';   // 把该处的字符替换为空字符
```
如果strchr()未找到换行符,fgets()在达到行末尾之前就达到了它能读取的最大字符数。可以像
在s_gets()中那样，给if添加一个else来处理。
