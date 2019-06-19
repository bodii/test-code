#include <stdio.h>
#include <sys/ipc.h>
#include <sys/msg.h>
#include <stdlib.h>
#include <string.h>
#include "queue.h"


void report_and_exit(const char* msg) {
	perror(msg);
	exit(-1);
}

int main() {
	key_t key = ftok(PathName, ProjectId);
	if (key < 0) report_and_exit("couldn't get key...");

	int qid = msgget(key, 0666 | IPC_CREAT);
	if (qid < 0) report_and_exit("couldn't get queue id...");

	char* payloads[] = {"msg1", "msg2", "msg3", "msg4", "msg5", "msg6"};
	int types[] = {1, 1, 2, 2, 3, 3};
	int i;
	for (i = 0; i < MsgCount; i++) {
		queuedMessage msg;
		msg.type = types[i];
		strcpy(msg.payload, payloads[i]);

		msgsnd(qid, &msg, sizeof(msg), IPC_NOWAIT);
		printf("%s sent as type %i\n", msg.payload, (int) msg.type);
	}

	return 0;
}

