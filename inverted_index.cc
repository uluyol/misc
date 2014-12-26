#include <fstream>
#include <iostream>
#include <map>
#include <sstream>
#include <string>
#include <vector>

bool is_whitespace(char c) {
	return c == ' ' || c == '\t' || c == '\n';
}

void add_words(std::vector<std::string> &dest, const std::string &s) {
	char *begin, *end;
	char t;
	end = begin = (char*)s.c_str();
	while (*begin != '\0') {
		while (!is_whitespace(*end) && *end != '\0')
				end++;
		if (*end == '\0') {
			dest.push_back(std::string(begin));
			break;
		}
		t = *end;
		*end = '\0';
		dest.push_back(std::string(begin));
		*end = t;
		for (begin = ++end; is_whitespace(*begin); begin = ++end);
	}
}

void populate_iindex(std::map<std::string, std::map<std::string, int>> &iindex, const std::string &fname) {
	std::ifstream in(fname);
	if (!in)
		throw in.rdstate();
	std::stringstream txt;
	txt << in.rdbuf();
	in.close();
	std::vector<std::string> words;
	add_words(words, txt.str());
	for (auto &w : words)
		iindex[w][fname] += 1;
}

int main() {
	std::map<std::string, std::map<std::string, int>> iindex;
	std::vector<std::string> files = {"bloom.cc", "cache.go", "quicksort.c", "self-repl.py"};
	for (auto &f : files)
		populate_iindex(iindex, f);
	std::string input;
	while (std::cin.good()) {
		std::cout << "Enter your search term: ";
		if (!std::getline(std::cin, input) || !iindex.count(input))
			goto endloop;
		for (auto &it : iindex[input])
			std::cout << it.first << ": " << it.second << "\n";
		endloop:
		std::cout << "\n";
	}
	return 0;
}