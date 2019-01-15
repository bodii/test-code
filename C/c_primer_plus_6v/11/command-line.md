#### 命令行参数
`命令行(command line)`是在命令行环境中，用户为运行程序输入命令的行.
`命令行参数(command-line argument)`是同一行的附加项。

'''c
/* repeat.c */
int main(int argc, char * argv[])
{
}
```
如运行时：
repeat    I'm     fine
argv[0]   argv[1] argv[2] 

argc = 3 3个字符串

C编译器允许main()没有参数或者两个参数(一些实现允许main()有更多参数，属于对标准的扩展)

main()有两个参数时，第1个参数是命令行中的字符串数量。过去，这个int类型的参数被称为argc
(表示参数计数(argument count)).系统用空格表示一个字符串的结束和下一个字符串的开始。
这个指向指针的指针称为argv(表示参数值(argument value)).
该程序本身的名称是argv[0],然后把随后的第1个字符串赋给argv[1],以此类推。

```c
int main(int argc, char **argv);

// 等价

int main(int argc, char *argv[]);
```
也就是说，argv是一个指向指针的指针，它所指向的指针指向char.因此，即使在原始定义中，argv
也是指向指针(该指针指向char)的指针。两种形式都可以使用。

许多环境(包括UNIX和DOS)都允许用双引号把多个单词括起来形成一个参数。如:
```shell
repeat "I am hungry" now
```
这行命令把字符串"I am hungry"赋给argv[1],把now赋给argv[2].
