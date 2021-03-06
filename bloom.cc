#include <functional>
#include <iostream>
#include <stdexcept>
#include <string>
#include <vector>

#include <stdint.h>

class Bitstring {
	uint8_t *bytes;
	size_t siz;
public:
	Bitstring(size_t siz);
	~Bitstring();
	size_t length();
	void set(size_t k);
	bool get(size_t k);
	std::string to_string();
};

Bitstring::Bitstring(size_t siz) {
	bytes = new uint8_t[siz]();
	this->siz = siz;
}

Bitstring::~Bitstring() { delete bytes; }
size_t Bitstring::length() { return siz; }

void Bitstring::set(size_t k) {
	size_t pos = k / 8;
	size_t bit = k % 8;
	if (pos >= siz)
		throw std::invalid_argument("Bit position is too large");
	bytes[pos] |= 1 << bit;
}

bool Bitstring::get(size_t k) {
	size_t pos = k / 8;
	size_t bit = k % 8;
	if (pos >= siz)
		throw std::invalid_argument("Bit position is too large");
	return (bytes[pos] >> bit) & 1;
}

std::string Bitstring::to_string() {
	std::string s(sizeof(uint8_t) * siz, '0');
	for (int i = 0; i < siz; i++) {
		for (int j = 0; j < sizeof(uint8_t); j++) {
			size_t bit = i * sizeof(uint8_t) + j;
			s[bit] = '0' + (char)this->get(bit);
		}
	}
	return s;
}

// Not as general as should be. We should allow many hashing functions, but we only support
// one at the moment
template <typename T>
class BloomFilter {
	Bitstring bits;
	void make_hashes(T obj, std::vector<size_t> &hashes);
public:
	BloomFilter(size_t siz);
	void add(T obj);
	bool has(T obj);
	std::string to_string();
};

template <typename T>
BloomFilter<T>::BloomFilter(size_t siz) : bits(siz) {}

template <typename T>
void BloomFilter<T>::make_hashes(T obj, std::vector<size_t> &hashes) {
	size_t shift_amount = sizeof(size_t) / 2;
	std::hash<T> hfunc;
	size_t hash = hfunc(obj);
	hashes.push_back((hash << shift_amount) >> shift_amount);
	hashes.push_back(hash >> shift_amount);
}

template <typename T>
void BloomFilter<T>::add(T obj) {
	std::vector<size_t> hashes;
	make_hashes(obj, hashes);
	for (auto &hash : hashes) {
		bits.set(hash % bits.length());
	}
}

template <typename T>
bool BloomFilter<T>::has(T obj) {
	std::vector<size_t> hashes;
	make_hashes(obj, hashes);
	bool has = true;
	for (auto &hash : hashes) {
		has &= bits.get(hash % bits.length());
	}
	return has;
}

template <typename T>
std::string BloomFilter<T>::to_string() {
	return bits.to_string();
}

int main(int argc, const char **argv) {
	BloomFilter<std::string> bf(10);
	std::cout << bf.to_string() << std::endl;
	bf.add("Happy");
	std::cout << bf.to_string() << std::endl;
	bf.add("Sad");
	std::cout << bf.to_string() << std::endl;
	std::cout << "Contains: \"Happy\": " << bf.has("Happy") << std::endl;
	std::cout << "Contains: \"Sad\": " << bf.has("Sad") << std::endl;
	std::cout << "Contains: \"Not sad\": " << bf.has("Not sad") << std::endl;
	return 0;
}