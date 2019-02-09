/* 定义一个宏函数,打印两个表达式及其值。例如,若参数为3+4和
4*12,则打印:3+4 is 7 and 4*12 is 48 */

#include <stdio.h>
#include <stdlib.h>

#define SUM(X, Y) (X+Y)
#define MUL(X, Y) (X*Y)
#define PR(X, Y, S, M) printf("%s is %d and %s is %d", X, S, Y, M)

struct exp {
	int x;
	char ex;
	int y;
};

int cal(struct exp *);
struct exp parse(char *);

int main(int argc, char * argv[])
{
	int a1, a2;
	struct exp b1,  b2;
	b1 = parse("3+4");
	b2 = parse("4*12");
	a1 = cal(&b1);
	a2 = cal(&b2);
	PR("3+4", a1, "4*12", a2);
	printf("%d\n", e.y); 

	return 0;
}

int cal(struct exp * e)
{
	int val;
	switch (e->ex)
	{
		case '+': val = SUM(e->x, e->y); break;
		case '*': val = MUL(e->x, e->y); break;
		default : val  = 0; break;
	}

	return val;
}

struct exp parse(char * str)
{
	struct exp e;
	e.x = str[0];
	e.ex = str[1];
	e.y = str[2];

	return e;
}
