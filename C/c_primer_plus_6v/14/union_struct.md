#### 联合(union)
联合是一种数据类型， 它能在同一个内存空间中储存不同的数据类型(不是同时储存)。
其典型的用法是，设计一种表以储存既无规律、事先也不知道顺序的混合类型。使用联
合类型的数组，其中的联合都大小相等，每个联合可以储存各种数据类型。

创建联合和创建结构的方式相同，需要一个联合模板和联合变量。
```c
// 一个带标记的联合模板
union hold {
	int digit;
	double bigfl;
	char letter;
};

// 定义3个与hold类型相关的变量
union hold fit;		  // hold类型的联合变量
union hold save[10;   // 内含10个联合变量的数组 
union hold * pu;      // 指向hold类型联合变量的指针
```
第1个声明创建了一个单独的联合变量fit。编译器分配足够的空间以便它能储存联合声明
中占用最大字节的类型。在本例中，占用空间最大的是double类型的数据。在我们的系统
中，double类型占64位，即8字节。第2个声明创建了一个数组save,内含10个元素，每个元
素都是8字节。第3个声明创建了一个指针，该指针变量储存hold类型联合变量的地址。

```c
// 初始化
union hold valA;
valA.letter = 'R';

union hold valB = valA;    // 用另一个联合来初始化
union hold valC = {88};    // 初始化联合的digit成员
union hold valD = {.bigfl = 118.2};  // 指定初始化器
```

// 使用联合
```c
fit.digit = 23;    // 把23储存在fit, 占2字节
fit.bigfl = 2.0;   // 清除23,储存2.0,占8字节
fit.letter = 'h';  // 清除2.0, 储存h, 占1字节
```
点运算符表示正在使用哪种数据类型。在联合中，一次只储存一个值。即使用有足够的空间，
也不能同时储存一个char类型值和一个int类型值。

和用指针访问结构使用->运算符一样，用指针访问联合时也要使用->运算符:
```c
pu = &fit;

x = pu->digit;    // 相当于x = fit.digit
```

联合的另一种用法是，在结构中储存与其成员有从属关系的信息。
```c
struct owner {
	char socsecurity[12];
};

struct leasecompany {
	char name[40];
	char headquarters[40];
};

union data {
	struct owner owncar;
	struct leasecompany leasecar;
}

struct car_data {
	char make[15];
	int status;  /* 私有为0, 租赁为1 */
	union data ownerinfo;
};
```
假设flits是car_data类型的结构变量，如果flits.status为0,程序将使用flits.ownerinfo.
owncar.socsecurity,如果flits.status为1,程序则使用flits.ownerinfo.leasecar.name。

#### 匿名联合(C11)
```c
struct owner {
	char socsecurity[12];
};
struct leasecompany {
	char name[40];
	char headquarters[40];
};
struct car_data {
	char make[15];
	int status;  /* 私有为0, 租赁为1 */
	union {
		struct owner owncar;
		struct leasecompany leasecar;
	}
};
```
现在，如果flits是car_data类型的结构变量，可以用flits.owncar.socsecurity 代替flits.
ownerinfo.owncar.socsecurity。

联合使用成员运算符的方式与结构相同：
```c
struct {
	int code;
	float cost;
} item;

item.code = 1265;
```

联合使用间接成员运算符的方式与结构相同：
```c
struct {
	int code;
	float cost;
} item, * ptrst;

ptrst = &item;
ptrst->code = 3451;
```
如下3个表达式是等价的：
ptrst->code    item.code  (*ptrst).code

