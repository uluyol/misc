#include <stdint.h>
#include <stdio.h>

#define BINSTRLEN sizeof(unsigned long)*8

static char _binstr[BINSTRLEN+1];

typedef struct {
	char vals[BINSTRLEN];
	int len;
} binStack;

int bsPush(binStack *st, char val) {
	if (st->len >= BINSTRLEN)
		return -1;

	st->vals[st->len++] = val;
	return 0;
}

int bsPop(binStack *st) {
	if (!st->len)
		return -1;

	return (int)st->vals[--st->len];
}

char *ulongToBin(unsigned long n) {
	static binStack bs;
	bs.len = 0;
	int pos = 0;
	int c;

	while (n != 0) {
		bsPush(&bs, (char)(n & 1) + '0');
		n >>= 1;
	}

	while ((c = bsPop(&bs)) != -1)
		_binstr[pos++] = (char)c;

	_binstr[pos] = '\0';

	return _binstr;
}

int32_t insertInto(int32_t M, int32_t N, int i, int j) {
	int32_t mask = ((1 << (j + 1 - i)) - 1) << i;
	return (N & ~mask) | (M<<i);
}

int main() {
	int32_t N = 1<<10;
	int32_t M = 19;
	int32_t result = insertInto(M, N, 2, 6);
	printf("Result: %s\n", ulongToBin(result));
	return 0;
}