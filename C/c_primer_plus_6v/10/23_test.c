#include <stdio.h>
#define ROWS 3
#define COLS 5

void copy2d(int, int, const double [][COLS], double [ROWS][COLS]);
void print2arr(int, int,const double [ROWS][COLS], const double [ROWS][COLS]);

int main(void)
{
	double ar1[ROWS][COLS] = {
		{ 1.2, 2.4, 5.7, 3.3, 2.8 },
		{ 3.6, 4.4, 7.9, 9.9, 1.5 },
		{ 5.2, 4.7, 9.1, 3.5, 6.8 },
	};

	double ar2[ROWS][COLS];

	copy2d(ROWS, COLS, ar1, ar2); 
	print2arr(ROWS, COLS,ar1, ar2);

	return 0;
}

void copy2d(int rows, int cols, const double ar1[rows][cols], double ar2[rows][cols])
{
	int r, c;
	for (r = 0; r < rows; r++)
		for (c = 0; c < cols; c++)
			//ar2[r][c] = ar1[r][c];
			// or
			ar2[r][c] = *( *(ar1 + r ) + c);
}

void print2arr(int rows, int cols, const double ar1[rows][cols], const double ar2[rows][cols])
{
	int r, c;
	for (r = 0; r < rows; r++)
		for (c = 0; c < cols; c++)
			printf("%.1f %.1f\n", ar1[r][c], ar2[r][c]);

}
