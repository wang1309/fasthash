DEBUG_OPT :=
OPT_LV := -O3
CFLAGS=-std=c99 -I./external/decNumber -g ${OPT_LV} -Wall -Werror
OBJS=mo.o arith.o compare.o logic.o decimal.o

all: libmo.a

libmo.a: $(OBJS)
	$(MAKE) ${DEBUG_OPT} -C external/decNumber
	ar -rcs libmo.a *.o

.PHONY: debug
debug: override OPT_LV := -O0
debug: override DEBUG_OPT := debug
debug: all

.PHONY: clean
clean:
	rm -f *.o *.a
