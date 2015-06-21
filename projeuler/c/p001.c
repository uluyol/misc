#include <stdio.h>

int main()
{
	int sum = 0;
	int i;

	for (i = 3; i < 1000; i += 3)
		sum += i;

	for (i = 5; i < 1000; i += 5)
		sum += (1 ^ !(i % 3)) * i;

	printf("Problem 1: %d\n", sum);
	return 0;
}