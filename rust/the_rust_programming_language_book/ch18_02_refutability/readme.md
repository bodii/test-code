### Refutability （可反驳性）:模式是否会匹配失效
模式有两种形式: `refutable（可反驳的）` 和 `irrefutable（不可反驳的）`。能匹配任何传递的可能值的模式被称为是`不可反驳的(irrefutable)`。

函数参数、let语句和for循环是只能授受不可反驳的模式，因为通过不匹配的值程序无法进行有意义的工作。
if let 和while let 表达式被限制为只能接受可反驳的模式，因为根据定义他们意在处理可能的失败：条件表达式的功能就是根据成功或失败执行不同的操作。

match匹配分支必须使用可反驳模式，除了最后一个分支需要使用能匹配任何剩余值的不可反驳模式。