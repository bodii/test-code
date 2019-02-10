#include <stdio.h>
#include <stdlib.h>

#define isBool(X) _Generic((X),\
	_Bool: "boolean",\
	default: "not boolean"\
		)


int main(int argc, char * argv[])
{
	int a = 1;
	printf("%s: %s\n", "a", isBool(a));
	return 0;
}
