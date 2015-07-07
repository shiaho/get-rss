ifeq ($(OS),Windows_NT)
  LIB_TARGET="libget_rss_win.a"
else
  LIB_TARGET="libget_rss.a"
endif

all:
	gcc -I. -c get_rss.c
	ar rvs $(LIB_TARGET) get_rss.o
	rm get_rss.o
