#### for 循环
> 关键字for 后面的圆括号中有3个表达式，分别用两个分号隔开。
第1个表达式是初始化，只会在for循环开始时执行一次。
第2个表达式是测试条件，在执行循环之前对表达式求值。
第3个表达式执行更新，在每次循环结束时求值。

for 循环的第1行包含了循环所需的所有信息：num的初值，num的终值和
每次循环num的增量。


#### for形式
`
for （ initialize; test; update )
	statement
`
在test 为假或0之前，重复执行statement部分
示例：
```c
for (n = 0; n < 10; n++)
	printf(" %d %d\n", n, 2 * n + 1);
```
