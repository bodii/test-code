#include <stdio.h>
#include "starfolk.h"


int main(void)
{
	struct bem * pb;

	struct bem deb = {
		6,
		{ "Berbnazel", "Gwolkapwolk" },
		"Arcturan"
	};

	pb = &deb;

	show(pb);

	return 0;
}


void show (struct bem * b)
{
	printf("%s %s is a %d-limbed %s.\n",
			b->title.first,
			b->title.last,
			b->limbs,
			b->type
		  );
}
