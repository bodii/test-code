struct name {
	char first[20];
	char last[20];
};

struct bem {
	int limbs;
	struct name title;
	char type[30];
};


void show(struct bem *);
