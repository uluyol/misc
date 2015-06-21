#include <glib.h>
#include <math.h>
#include <stdint.h>
#include <stdlib.h>
#include "libprojeuler.h"

// Need to free after use with g_array_free. Returns array of factors.
GArray *get_prime_factors(int64_t n)
{
	guint8 *sieve;
	GArray *factors = g_array_new(TRUE, FALSE, sizeof(int64_t));
	int64_t i, j, sieve_len;
	int64_t two = 2;

	sieve_len = ceil(sqrt(n))+1;
	sieve = calloc(sieve_len, sizeof(guint8));

	if (n % 2 == 0)
		g_array_append_val(factors, two);

	for (i = 3; i < sieve_len; i+=2) {
		if (sieve[i])
			continue;
		for (j = i; j < sieve_len; j += i)
			sieve[j] = TRUE;
		if (n % i == 0)
			g_array_append_val(factors, i);
	}

	free(sieve);
	return factors;
}