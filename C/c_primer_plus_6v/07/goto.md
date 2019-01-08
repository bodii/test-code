#### goto语句
goto语句有两部分：goto和标签名。标签的命名遵循变量命名规则，如
```c
goto part2;
```
要让这条语句正常工作，函数还必须包含另一条标为part2的语句，该语句
以标签名后紧跟一个冒号开始:
```c
part2: printf("Refined analysis:\n");
```


#### 避免使用goto
原则上，根本不用在C程序中使用goto语句。
`建议谨慎使用，或者根本不用`
