
SRCS = $(wildcard p*.go | sort)
OBJS = $(SRCS:.go=)
LIBS = libprojeuler.go

all: $(OBJS)

%: %.go
	@echo "  GO  " $<
	@go build $< $(LIBS)

clean:
	rm -f $(OBJS)