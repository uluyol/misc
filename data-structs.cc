#include <iostream>
#include <stdexcept>

// Stack

template <typename T>
struct Node {
	T val;
	Node<T> *next;
};

template <typename T>
class Stack {
	Node<T> *top;
public:
	Stack();
	~Stack();
	void push(T val);
	T pop();
	T peek();
	bool is_empty();
};

template <typename T>
Stack<T>::Stack() {
	top = nullptr;
}

template <typename T>
Stack<T>::~Stack() {
	while (!is_empty())
		pop();
}

template <typename T>
void Stack<T>::push(T val) {
	Node<T> *n = new Node<T>;
	n->val = val;
	n->next = top;
	top = n;
}

template <typename T>
T Stack<T>::pop() {
	T val = peek();
	Node<T> *cur = top;
	top = top->next;
	delete cur;
	return val;
}

template <typename T>
T Stack<T>::peek() {
	if (is_empty()) {
		throw std::out_of_range("Stack is empty");
	}
	return top->val;
}

template <typename T>
bool Stack<T>::is_empty() {
	if (top == nullptr)
		return true;
	return false;
}

// Hash Table
#define MAP_MAXLOADFACTOR 0.75

template <typename T>
struct Entry {
	std::string key;
	T val;
};

// Toy hashing function
static size_t makeHash(std::string key) {
	size_t hash = 4226724;
	size_t pos = 0;
	for (char &c : key) {
		hash ^= c << pos;
		pos = (pos + 1) % sizeof(size_t);
	}
	return hash;
}

template <typename T>
static int insertEntry(Node<Entry<T>> **slots, size_t cap, Entry<T> entry) {
	Node<Entry<T>> *cur;
	size_t pos = makeHash(entry.key) % cap;
	for (cur = slots[pos]; cur != nullptr; cur = cur->next) {
		if (entry.key == cur->val.key) {
			cur->val = entry;
			return 0;
		}
	}
	cur = new Node<Entry<T>>;
	cur->val = entry;
	cur->next = slots[pos];
	slots[pos] = cur;
	return 1;
}

template <typename T>
static void insertNodeWithoutSearch(Node<Entry<T>> **slots, size_t cap, Node<Entry<T>> *node) {
	size_t pos = makeHash(node->val.key) % cap;
	node->next = slots[pos];
	slots[pos] = node;
}

template <typename T>
class Map {
	Node<Entry<T>> **slots;
	size_t cap;
	size_t entries;
	void grow();
public:
	Map();
	~Map();
	void set(std::string key, T val);
	T get(std::string key);
	void remove(std::string key);
};

template <typename T>
Map<T>::Map() {
	cap = 16;
	entries = 0;
	slots = new Node<Entry<T>>*[cap]();
}

template <typename T>
Map<T>::~Map() {
	for (size_t i = 0; i < cap; i++) {
		if (slots[i] == nullptr)
			continue;
		Node<Entry<T>> *prev, *cur;
		prev = slots[i];
		for (cur = prev->next; cur != nullptr; cur = cur->next) {
			delete prev;
			prev = cur;
		}
		delete prev;
	}
	delete slots;
}

template <typename T>
void Map<T>::grow() {
	size_t newcap = cap * 2;
	Node<Entry<T>> *cur, *next;
	Node<Entry<T>> **new_slots = new Node<Entry<T>>*[newcap]();
	for (size_t i = 0; i < cap; i++) {
		for (cur = slots[i]; cur != nullptr; cur = next) {
			next = cur->next;
			insertNodeWithoutSearch(new_slots, newcap, cur);
		}
	}
	delete slots;
	slots = new_slots;
	cap = newcap;
}

template <typename T>
void Map<T>::set(std::string key, T val) {
	if (((float)entries / (float)cap) >= MAP_MAXLOADFACTOR)
		grow();

	Entry<T> entry = {key, val};
	entries += insertEntry(slots, cap, entry);
}

template <typename T>
T Map<T>::get(std::string key) {
	Node<Entry<T>> *cur;
	size_t pos = makeHash(key) % cap;
	for (cur = slots[pos]; cur != nullptr; cur = cur->next) {
		if (key == cur->val.key)
			return cur->val.val;
	}
	
	throw std::out_of_range("Key not found");
}

template <typename T>
void Map<T>::remove(std::string key) {
	Node<Entry<T>> *cur;
	Node<Entry<T>> *prev = nullptr;
	size_t pos = makeHash(key) % cap;
	for (cur = slots[pos]; cur != nullptr; cur = cur->next) {
		if (key == cur->val.key) {
			if (prev == nullptr)
				slots[pos] = cur->next;
			else
				prev->next = cur->next;
			delete cur;
			return;
		}
		prev = cur;
	}
}

int main() {
	Stack<int> st;

	/* Stack test */
	st.push(1);
	st.push(2);
	st.push(3);
	st.push(4);
	st.push(5);
	st.push(6);
	st.push(7);

	while (!st.is_empty())
		std::cout << st.pop() << std::endl;

	/* Hashmap test */

	Map<int> map;
	map.set("1>", 1);
	map.set("2>", 2);
	map.set("3>", 3);
	map.set("4>", 4);
	map.set("5>", 5);
	map.set("6>", 6);
	map.set("7>", 7);
	map.set("8>", 8);
	map.set("9>", 9);

	std::cout << map.get("1>") << std::endl;
	std::cout << map.get("2>") << std::endl;
	std::cout << map.get("3>") << std::endl;
	std::cout << map.get("4>") << std::endl;
	std::cout << map.get("5>") << std::endl;
	std::cout << map.get("6>") << std::endl;
	std::cout << map.get("7>") << std::endl;
	std::cout << map.get("8>") << std::endl;
	std::cout << map.get("9>") << std::endl;

	map.remove("1>");
	try {
		map.get("1>");
		std::cout << "Failed: key '1>' was removed\n";
	} catch (const std::out_of_range& e) {
		std::cout << "Success: key: '1>' was removed\n";
	}
	map.set("2>", 22);
	std::cout << map.get("2>") << std::endl;
	map.remove("2>");
	try {
		map.get("2>");
		std::cout << "Failed: key '2>' was removed\n";
	} catch (const std::out_of_range& e) {
		std::cout << "Success: key: '2>' was removed\n";
	}

	return 0;
}
