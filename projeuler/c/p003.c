#include <glib.h>
#include <inttypes.h>
#include <stdint.h>
#include <stdio.h>
#include "libprojeuler.h"

int main()
{
	int64_t n = 600851475143;
	int64_t cur;
	int64_t biggest;
	int i;
	GArray *factors;

	factors = get_prime_factors(n);

	for (i = 0;; i++) {
		cur = g_array_index(factors, int64_t, i);

		if (!cur)
			break;

		biggest = MAX(biggest, cur);
	}

	printf("Problem 3: %" PRId64 "\n", biggest);

	g_array_unref(factors);
	return 0;
}