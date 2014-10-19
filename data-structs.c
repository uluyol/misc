#include <stdio.h>
#include <stdlib.h>
#include <string.h>

/* Stack */

typedef struct Node Node;
struct Node {
	int val;
	Node *next;
};

// This only takes non-negative values
typedef struct {
	Node *top;
} Stack;

int StackPush(Stack *st, int val) {
	Node *top = malloc(sizeof(Node));
	if (top == NULL)
		return -1;
	top->val = val;
	top->next = st->top;
	st->top = top;
	return 0;
}

int StackPop(Stack *st) {
	int val;
	Node *np = st->top;

	if (np == NULL)
		return -1;
	val = np->val;
	st->top = np->next;
	free(np);
	return val;
}

int StackPeek(Stack *st) {
	if (st->top == NULL)
		return -1;
	return st->top->val;
}

/* Hash Table */

#define MAXLOADFACTOR 0.75

typedef struct MapEntry MapEntry;
struct MapEntry {
	char *key;
	int val;
	MapEntry *next;
};

typedef struct {
	MapEntry **slots;
	int cap;
	int entries;
} Map;

int MapInit(Map *map) {
	map->cap = 16;
	map->slots = calloc(map->cap, sizeof(MapEntry*));
	if (map->slots == NULL)
		return -1;
	map->entries = 0;
	return 0;
}

static int makeHash(char *key) {
	int hash = 4226724;
	int pos = 0;
	while (*key != '\0') {
		hash ^= *key << pos;
		pos++;
		key++;
	}
	return hash >= 0 ? hash : ~hash;
}

static int updateMapEntry(MapEntry **slots, int cap, MapEntry *me) {
	MapEntry *cur;
	int pos = makeHash(me->key) % cap;
	for (cur = slots[pos]; cur != NULL; cur = cur->next) {
		if (strcmp(me->key, cur->key) == 0) {
			cur->val = me->val;
			return 0;
		}
	}
	me->next = slots[pos];
	slots[pos] = me;
	return 1;
}

static void insertMapEntry(MapEntry **slots, int cap, MapEntry *me) {
	int pos = makeHash(me->key) % cap;
	me->next = slots[pos];
	slots[pos] = me;
}

static void mapGrow(Map *map) {
	int i, newcap = map->cap*2;
	MapEntry *cur, *next;
	MapEntry **newSlots = calloc(newcap, sizeof(MapEntry*));
	if (newSlots == NULL)
		return;
	for (i = 0; i < map->cap; i++) {
		for (cur = map->slots[i]; cur != NULL; cur = next) {
			next = cur->next;
			insertMapEntry(newSlots, newcap, cur);
		}
	}
	free(map->slots);
	map->slots = newSlots;
	map->cap = newcap;
}

int MapInsert(Map *map, char *key, int val) {
	if ((map->entries / map->cap) >= MAXLOADFACTOR)
		mapGrow(map);

	MapEntry *me = malloc(sizeof(MapEntry));
	if (me == NULL)
		return -1;
	me->key = key;
	me->val = val;
	map->entries += updateMapEntry(map->slots, map->cap, me);

	return 0;
}

int MapGet(Map *map, char *key) {
	MapEntry *cur;
	int pos = makeHash(key) % map->cap;
	for (cur = map->slots[pos]; cur != NULL; cur = cur->next) {
		if (strcmp(cur->key, key) == 0)
			return cur->val;
	}

	return -1;
}

void MapRemove(Map *map, char *key) {
	MapEntry *prev = NULL, *cur;
	int pos = makeHash(key) % map->cap;
	for (cur = map->slots[pos]; cur != NULL; cur = cur->next) {
		if (strcmp(cur->key, key) == 0) {
			if (prev == NULL)
				map->slots[pos] = cur->next;
			else
				prev->next = cur->next;
			free(cur);
			return;
		}
		prev = cur;
	}
}

int main() {
	Stack st;
	int val;
	st.top = NULL;

	/* Stack test */

	StackPush(&st, 1);
	StackPush(&st, 2);
	StackPush(&st, 3);
	StackPush(&st, 4);
	StackPush(&st, 5);
	StackPush(&st, 6);
	StackPush(&st, 7);

	while ((val = StackPop(&st)) != -1)
		printf("%d\n", val);

	/* Hashmap test */

	Map map;
	MapInit(&map);
	MapInsert(&map, "1>", 1);
	MapInsert(&map, "2>", 2);
	MapInsert(&map, "3>", 3);
	MapInsert(&map, "4>", 4);
	MapInsert(&map, "5>", 5);
	MapInsert(&map, "6>", 6);
	MapInsert(&map, "7>", 7);
	MapInsert(&map, "8>", 8);
	MapInsert(&map, "9>", 9);

	printf("%d\n", MapGet(&map, "1>"));
	printf("%d\n", MapGet(&map, "2>"));
	printf("%d\n", MapGet(&map, "3>"));
	printf("%d\n", MapGet(&map, "4>"));
	printf("%d\n", MapGet(&map, "5>"));
	printf("%d\n", MapGet(&map, "6>"));
	printf("%d\n", MapGet(&map, "7>"));
	printf("%d\n", MapGet(&map, "8>"));
	printf("%d\n", MapGet(&map, "9>"));

	MapRemove(&map, "1>");
	printf("%d\n", MapGet(&map, "1>"));

	MapInsert(&map, "2>", 22);
	printf("%d\n", MapGet(&map, "2>"));
	MapRemove(&map, "2>");
	printf("%d\n", MapGet(&map, "2>"));

	return 0;
}