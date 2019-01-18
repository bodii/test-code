#### 其他标准I/O函数

#### int ungetc(int c, FILE *fp)函数
int ungetc()函数把C指定的字符放回输入流中。如果把一个字符放回输入流，
下次调用标准输入函数时将读取该字符。

例如，假设要读取下一个冒号之前的的所有字符，但是不包括冒号本身，可以
使用getchar()或getc()函数读取字符到冒号，然后使用ungetc()函数把冒号放
回输入流中。ANIS C标准保证每次只会放回一个字符。如果实现允许把一行中的
多个字符放回输入流，那么下一次输入函数读入的字符顺序与放回时的顺序相反。


#### int fflush()函数
```c
int fflush(FILE *fp);
```
调用fflush()函数引起输出缓冲区中所有的未写入数据被发送到fp指定的输出文件。
这个过程称为刷新缓冲区。如果fp是空指针，所有输出缓冲区都被刷新。在输入流中
使用fflush()函数的效果是未定义的。只要最近一次操作不是输入操作，就可以用该
函数来更新流(任何读写模式)


#### int setvbuf()函数
```c
int setvbuf(FILE * restrict fp, char * restrict buf, int mode, size_t size);
```
setvbuf()函数创建了一个供标准I/O函数替换使用的缓冲区。在打开文件后且未对流进
行其他操作之前，调用该函数。指针fp识别待处理的流,buf指向待使用的存储区。如果
buf的值不是NULL，则必须创建一个缓冲区。
例如，声明一个内含1024个字符的数组，并传递该数组的地址。然而，如果把NULL作为
buf的值，该函数会为自己分配一个缓冲区。变量size告诉setvbuf()数组的大小(size_t
是一种派生的整数类型).mode的选择如下：_IOFBF表示完全缓冲(在缓冲区满时刷新);
_IOLBF表示行缓冲(在缓冲区满时或写入一个换行符时);_IONBF表示无缓冲。
如果操作成功，函数返回0,否则返回一个非零值。

假设一个程序要储存一种数据对象，每个数据对象的大小是3000字节。可以使用setvbuf()
函数创建一个缓冲区，其大小是该数据对象大小的倍数。


#### 二进制I/O：frea()和fwrite()
如果以程序所用的表示法把数据储存在文件中，则称以二进制形式储存数据。
不存在从数值形式到字符串的转换过程。对于标准I/O，fread()和fwrite函数用于以二进
制形式处理数据。
实际上，所有的数据都是以二进制形式储存的，甚至连字符都以字符码的二进制表示来储存
。如果文件中的所有数据都被解释成字符码，则该文件包含文本数据。如果部分或所有的数
据都被解释成二进制的数值数据，则称该文件包含二进制数据。

可以调用getc()拷贝包含二进制数据的文件。
一般而言，用二进制模式在二进制格式文件中储存二进制数据。


#### size_t fwrite()函数
```c
size_t fwrite(const void * restrict ptr, size_t size, size_t nmemb, FILE * restrict fp);
```
fwrite()函数把二进制数据写入文件。
```c
char buffer[256];
fwrite(buffer, 256, 1, fp);
```
```c
double earnings[10];
fwrite(earnings, sizeof(double), 10, fp);
```


#### size_t fread()函数
```c
size_t fread(void * restrict ptr, size_t size, size_t nmemb, FILE * restrict fp);
```
```c
double earnings[10];
fread(earnings, sizeof(double), 10, fp);
```


#### feof()和ferror()函数
```c
int feof(FILE *fp);
int ferror(FILE *fp);
```
如果标准输入函数返回EOF,则通常表示函数已到达文件结尾。然而，出现读取错误时，函数也会
返回EOF.feof()和ferror()函数用于区分这两种情况。
当上一次输入调用检测到文件结尾时，feof()函数返回一个非零值，否则返回0.
当读或写出错时,ferror()函数返回一个非零值，否则返回0.
