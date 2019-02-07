// 内联函数 放入头文件中

#ifndef EATLINE_H_
#define EATLINE_H_
inline static void eatline()
{
	while (getchar() != '\n')
		continue;
}
#endif
