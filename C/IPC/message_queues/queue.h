#define ProjectId  123
#define PathName   "queue.h"
#define MsgLen     4
#define MsgCount   6

typedef struct {
	long type;
	char payload[MsgLen + 1];
} queuedMessage;
