#### 与文件进行通信

#### 文件是什么
文件(file)通常烛在磁盘或固态硬盘上的一段已命名的存储区。
C把文件看作是一系列连续的字节，每个字节都能单独读取。这与UNIX环境中(C的发源地)
的文件结构相对应。由于其他环境中可能无法完全对应这个模型,C提供两种文件模式：
文本模式和二进制模式。


#### 文本模式和二进制模式
虽然C提供了二进制模式和文本模式，但是这两种模式的实现可以相同。前面提到过，因为
UNIX使用一种文件模式，这两种模式对于UNIX实现而言完全相同。Linux也是如此。


#### I/O的级别
I/O的两个级别，即处理文件访问的的两个级别.
底层I/O(low-level I/O)使用操作系统提供的基本I/O服务。
标准高级I/O(standard high-level I/O)使用C库的标准包和stdio.h头文件定义.

无法保证所有的操作系统都使用相同的底层I/O模型，C标准只支持标准I/O包。


#### 标准文件
C程序会自动打开3个文件，它们被称为`标准输入(standard input)`、`标准输出(standard 
output)`和`标准错误输出(standard error output)`.

在默认情况下，标准输入是系统的普通输入设备，通常为键盘;标准输出和标准错误输出是系
统的普通输出设备，通常为显示屏。


#### 标准I/O
与底层I/O相比，标准I/O包除了可移植以外还有两个好处。
第一，标准I/O有许多专门的函数简化了处理不同I/O的问题。
第二，输入和输出都是缓冲的。


#### getc()和putc()函数
getc()和putc()函数与getchar()和putchar()函数类似。所不同的是，要告诉getc()和putc
()函数使用哪一个文件：
```c
ch = getchar();
// 然而，"从fp指定的文件中获取一个字符":
ch = getc(fp);
// 把字符ch放入FILE指针fpout指定的文件中:
putc(ch, fpout);
```


#### 文件结尾
如果getc()函数在读取一个字符时发现是文件结尾，它将返回一个特殊值EOF.所以C程序只有在
读到超过文件末尾时才会发现文件的结尾。


#### fclose()函数
fclose(fp)函数关闭fp指定的文件，必要时刷新缓冲区。
如果成功关闭，fclose()函数返回0,否则返回EOF:
```c
if (fclose(fp) != 0)
	printf("Error in closing file %s\n", argv[1]);
```
如果磁盘已満，移动硬盘被移除或出现I/O错误，都有会导致调用fclose()函数失败。


#### 指向标准文件的指针

| 标准文件 | 文件指针 | 通常使用的设备 |
|-|:-|:-:|:-:|-:|
| 标准输入 | stdin | 键盘 |
| 标准输出 | stdout | 显示器 |
| 标准错误 | stderr | 显示器 |


#### fprintf()
fprintf和printf()类似，但是fprintf()的第1个参数必须是一个文件指针。程序中使用stderr
指针把错误消息发送至标准错误。

同时打开的文件数量是有限的，这个限制取决于系统和实现，范围一般是10~20.相同的文件指针
可处理不同的文件，前提是这些文件不需要同时打开。


#### 文件I/O： fprintf()、fscanf()、fgets()和fputs()
fprintf()和fscanf()的工作方式与printf()和scanf()类似。但是，与putc()不同的是，fprintf()
和fscanf()函数都把FILE指针作为第1个参数，而不是最后一个参数。
rewind()函数让程序回到文件开始处，方便while循环打印整个文件的内容。


#### fgets()和fputs函数
fgets()的第1个参数和gets()函数一样，也是表示储存输入位置的地址(char * 类型);
第2个参数是一个整数，表示待输入字符串的大小。
最后一个参数是文件指针，指定待读取的文件。
```c
fgets(buf, STLEN, fp);
```
这里，buf是char类型数组的名称，STLEN是字符串的大小，fp是指向FILE的指针。
fgets()函数读取输入直到第1个换行符的后面，或读到文件结尾，或者读取STLEN-1个字符。
然后，fgets()在末尾添加一个空字符使之成为一个字符串。
如果fgets()在读到字符上限之前已读完一整行，它会把表示结尾的换行符放在空字符前面。
fgets()函数在遇到EOF时返回NULL值，可以利用这一机制检查是否到达文件结尾;如果未遇到EOF则
之前返回传给它的地址。

fputs()函数接受两个参数：第1个是字符串的地址;
第2个是文件指针。
该函数根据传入地址找到的字符串写入指定的文件中。和puts()函数不同，
fputs()在打印字符串时不会在其末尾添加换行符。
```c
fputs(buf, fp);
```
这里的buf是字符串的地址，fp是用于指定目标文件。
fgets()保留了换行符，fputs()就不会再添加换行符。


#### 随机访问：fseek()和ftell()
fseek()函数，可把文件看作是数组，在fopen()打开的文件中直接移动到任意字节处。
fseek()有3个参数，返回int类型的值;
ftell()函数返回一个long类型的值，表示文件中的当前位置。


#### fseek()和ftell的工作原理
fseek()的第1个参数是FILE指针，指向待查找的文件，fopen应该已经打开该文件。
fseek()的第2个参数是偏移量(offset).该参数表示从起始点开始要移动的距离。该参数必须是一个long类型
的值，可以为正(前移)、负(后移)或0(保持不动);
fseek()的第3个参数是模式，该参数确定起始点。
根据ANSI标准，在stdio.h头文件中规定了几个表示模式的明示常量(manifest constant),如：
| 模式 | 偏移量的起始点 |
|-|:-:|-:|
| SEEK_SET | 文件开始处 |
| SEEk_CUR | 当前位置 |
| SEEK_END | 文件末尾 |

下面是调用fseek()函数的一些示例:
```c
fseek(fp, 0L, SEEK_SET);  // 定位至文件开始处
fseek(fp, 10L, SEEK_SET); // 定位至文件中的第10个字节
fseek(fp, 2L, SEEK_CUR);  // 从文件当前位置前移2个字节
fseek(fp, 0L, SEEK_END);  // 定位至文件结尾
fseek(fp, -10L, SEEK_END);// 从文件结尾处回退10个字节
```
如果一切正常，fseek()返回值为0;如果出现错误,其返回值是-1;


ftell()函数的返回类型是long,它返回的是当前的位置。
ANSI C把它定义在stdio.h中。
在最初实现的UNIX中，ftell()通过返回距文件开始处的字节数来确定文件的位置。文件的第1个字节
到文件开始处的距离是0，以此类推。ANSI C规定，该定义适用于以二进制模式打开的文件，以文件模
式打开文件的情况不同。
```c
fseek(fp, 0L, SEEK_END);
//把当前位置设置为距文件末尾0字节偏移量
last = fteel(fp);
//把从文件开始处到文件结尾的字节数赋给last.
// 然后是一个for循环:
for(count = 1L; count <= last; count++)
{
	fseek(fp, -count, SEEK_END);
	ch = getc(fp);
}
```


#### 二进制模式和文本模式
ftell()函数在文本模式和二进制模式中的工作方式不同。
许多系统的文本格式与UNIX的模型有很大不同，导致从文件开始处统计的字符成为一个毫无意义的值。


#### 可移植性
理论上， fseek()和ftell()应该符合UNIX模型。但是，不同系统存在着差异，有时确实无法做到与UNIX
模型一致。因此，ANSI对之些函数降低了要求。
在文本模式中，只有以下调用能保证其相应的行为:
| 函数调用 | 效果 |
|-|:-:|-:|
| fseek(file, 0L, SEEK_SET) | 定位到文件开始处 |
| fseek(file, 0L, SEEk_CUR) | 保持当前位置不动 |
| fseek(file, 0L, SEEK_END) | 定位至文件末尾 |
| fseek(file, ftell-pos, SEEK_set) | 到距文件开始处ftell-pos的位置，ftell-pos是ftell()的返回值 |


#### fgetpos()和fsetpos()函数
fseek()和ftell()潜在的问题是，它们都把文件大小限制在long类型能表示的范围内。
ANSI C新增了两个处理较大文件的新定位函数：fgetpos()和fsetpos().
这两个函数不使用long类型的值表示位置，它们使用一种新类型:fpos_t(代表file position type,文件定位
类型).
fpos_t可以实现为结构。
```c
int fsetpos(FILE *stream, const fpos_t *pos);
```
调用该函数时， 使用pos位置上的fpos_t类型值来设置文件指针指向该值指定的位置。如果成功，fsetpos()
函数返回0;如果失败，则返回非0.
fpos_t类型的值应通过之前调用fgetpos()获得。
