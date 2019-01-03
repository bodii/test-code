#### _Bool类型

C99标准添加了_Bool类型，用于表示布尔值，即逻辑值true和false。
因为c语言用值1表示true, 值0表示false, 所以_Bool类型实际上也是
一种整数类型。但原则上它仅占用1位存储空间，因为对0和1而言，1位
的存储空间足够了。