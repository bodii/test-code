// 千万不要模仿！

#include <stdio.h>

int main(void)
{
	char side_a[] = "Side A";
	char dont[] = { 'W', 'O', 'W', '!' };
	char side_b[] = "Side B";

	puts(dont); /* dont 不是一个字符串 */

	return 0;
}

/*
 由于dont缺少一个表示结束的空字符，所以它不是一个字符串，因此puts()不知道
 在何处停止。它会一直打印dont后面内存中的内容，直到发现一个空字符为止。为
 了让puts()能尽快读到空字符，我们把dont放在side_a和side_b之间。
 所以puts()一直输出至遇到side_a中的空字符。
*/
