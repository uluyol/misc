#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

/* From scratch implementation of quicksort to learn. Only handles ints for
 * simplicity
 */

void print_array(int *arr, int len) {
	int i;
	if (len == 0)
		return;
	for (i = 0; i < len-1; i++)
		printf("%d ", arr[i]);
	printf("%d\n", arr[len-1]);
}


void _swap(int *arr, int i, int j) {
	int t = arr[i];
	arr[i] = arr[j];
	arr[j] = t;
}

int _partition(int *arr, int first, int last, int pivot) {
	int pval, smalli, i;
	pval = arr[pivot];
	_swap(arr, pivot, last);
	smalli = first;
	for (i = first; i < last; i++) {
		if (arr[i] <= pval) {
			_swap(arr, i, smalli);
			smalli++;
		}
	}
	_swap(arr, smalli, last);
	return smalli;
}

void _qsort(int *arr, int first, int last) {
	int pivot;
	if (first >= last)
		return;

	pivot = first + ( rand() % ( last + 1 - first ) );
	pivot = _partition(arr, first, last, pivot);
	_qsort(arr, first, pivot-1);
	_qsort(arr, pivot+1, last);
}

void quicksort(int *arr, int len) {
	srand(time(NULL));
	_qsort(arr, 0, len-1);
}

int main() {
	int vals[] = {6, 1, 1, 7, 9, 1, 6, 9, 0, 2, 4, 7, 4, 5, 9, 6, 8, 3};
	int vals_len = sizeof(vals)/sizeof(int);

	printf("original: "); print_array(vals, vals_len);
	quicksort(vals, vals_len);
	printf("  sorted: "); print_array(vals, vals_len);
	return 0;
}