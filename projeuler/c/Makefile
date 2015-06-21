CC = clang
SRCS = $(wildcard p*.c | sort)
OBJS = $(SRCS:.c=)
LIBS = libprojeuler.c
FLAGS = -lm $(shell pkg-config --cflags --libs glib-2.0)

all: $(OBJS)

%: %.c
	@echo "  CC  " $<
	@$(CC) -Wall $(FLAGS) -std=c99 $(LIBS) $< -o $@

clean:
	rm -f $(OBJS)