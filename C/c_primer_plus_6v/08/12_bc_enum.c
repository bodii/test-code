// 一个提供加法、减法、乘法、除法的菜单
#include <stdio.h>
#include <ctype.h>
#include <stdbool.h>

void printfEnum(void); // 打印菜单

char checkEnteredEnumOptionVal(void); // 检测输入的菜单选项

float checkEnteredNumber(int flag); // 检测输入的数字(float);

void addition(void); // 加法

void subtraction(void); // 减法

void multiplication(void); // 乘法

void division(void); // 除法


int main(void)
{
	char ch;
	bool stop = false;

	while (!stop)
	{
		printfEnum();
		ch = checkEnteredEnumOptionVal();
	
		switch (ch)
		{
			case 'a': addition(); 
					  break;
			case 's': subtraction(); 
					  break;
			case 'm': multiplication(); 
					  break;
			case 'd': division(); 
					  break;
			case 'q': printf("Bye.\n");
					  stop = true;
					  break;
		}
	}

	return 0;
}


// 打印菜单
void printfEnum(void)
{
	printf("Enter the operation of your choice:\n");
	printf("%-10s\t\t%-10s\n"
		   "%-10s\t\t%-10s\n"
		   "%-10s\n",
		   "a. add", "s. subtract",
		   "m. multiply", "d. divide",
		   "q. quit"
		  );
}


// 检测输入的菜单选项
char checkEnteredEnumOptionVal(void)
{
	char ch;

	while ((ch = getchar()) != '\n') // 获取输入的字符
	{
		
		// 检测输入字符是否错误
		if (ch != 'a' && ch != 's' && ch != 'm' && ch != 'd' && ch != 'q')
		{
			printf("Input format error!\n");
			printfEnum();
		} 
		else if(ch == 'a' || ch == 's' || ch == 'm' || ch == 'd' || ch == 'q') 
		{
			break; // 跳出循环
		}

		// 过滤掉多余的输入
		while (getchar() != '\n')
			continue;

	}

	return ch;
}


// 检测输入的数字(float)
float checkEnteredNumber(int flag)
{
	float f;

	if (flag == 1)
		printf("Enter first number:");
	else
		printf("Enter second number:");

	while ((scanf("%f", &f) != 1 || (int)f == 0))
	{
		// 如果输入的不是数字
		if (getchar() != '0')
		{
			printf("one is not an number.\n");
			printf("Please enter a number, such as 2.5, -1.78E8, or 3: ");

			while (getchar() != '\n')
				continue;
		}
		else if((int)f == 0) // 如输入的是0 
		{
			printf("Enter a number other than 0:");

			while (getchar() != '\n')
				continue;
		}

	}

	return f;
}


// 加法
void addition() 
{
	float a, b;
	
	a = checkEnteredNumber(1); 
	b = checkEnteredNumber(2); 

	printf("%.2f + %.2f = %.2lf\n", a, b, (double)(a + b));
}


// 减法
void subtraction()
{
	float a, b;

	a = checkEnteredNumber(1); 
	b = checkEnteredNumber(2); 

	printf("%.2f - %.2f = %.2lf\n", a, b, (double)(a - b));
}


// 乘法
void multiplication()
{
	float a, b;

	a = checkEnteredNumber(1); 
	b = checkEnteredNumber(2); 

	printf("%.2f * %.2f = %.2lf\n", a, b, (double)(a * b));
}


// 除法
void division()
{
	float a, b;

	a = checkEnteredNumber(1); 
	b = checkEnteredNumber(2); 

	printf("%.2f / %.2f = %.2lf\n", a, b, (double)(a / b));
}

