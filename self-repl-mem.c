char s[] = "#include <stdio.h>\n#include <string.h>\nint main() {\nint i; printf(\"char s[] = {\"); for (i = 0; i < strlen(s); i++) printf(\"%d,\", s[i]); puts(\"0};\"); printf(\"int gen = %d;\", gen+1); puts(\"\");printf(\"%s\", s); return 1;}";
int gen = -1;

#include <stdio.h>
#include <string.h>

int main() {
	int i;
	printf("char s[] = {");
	for (i = 0; i < strlen(s); i++) printf("%d,", s[i]);
	puts("0};");
	printf("int gen = %d;", gen+1);
	puts("");
	printf("%s", s);
	return 1;
}