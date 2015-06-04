all:
	gcc `go env GOGCCFLAGS` -I. -c get_rss.c
	ar rvs libget_rss.a get_rss.o
	rm get_rss.o
