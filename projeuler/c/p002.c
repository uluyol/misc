#include <stdio.h>

int main()
{
	int prev = 1;
	int cur = 2;
	int sum = 0;

	while (cur < 4000000) {
		sum += cur;

		prev += 2 * cur;
		cur = 2 * prev - cur;
	}

	printf("Problem 2: %d\n", sum);
	return 0;
}