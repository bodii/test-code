#include <sys/wait.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>

#define ReadEnd   0
#define WriteEnd  1

void report_and_exit(const char* msg) {
	perror(msg);
	exit(-1);
}

int main() {
	int pipeFDs[2];
	char buf;
	const char* msg = "Nature's first green is gold\n";

	if (pipe(pipeFDs) < 0) report_and_exit("pipeFD");
	pid_t cpid = fork();
	if (cpid < 0) report_and_exit("fork");

	if (0 == cpid) {
		close(pipeFDs[WriteEnd]);

		while (read(pipeFDs[ReadEnd], &buf, 1) > 0)
			write(STDOUT_FILENO, &buf, sizeof(buf));

		close(pipeFDs[ReadEnd]);
		_exit(0);
	}
	else {
		close(pipeFDs[ReadEnd]);

		write(pipeFDs[WriteEnd], msg, strlen(msg));
		close(pipeFDs[WriteEnd]);

		wait(NULL);
		exit(0);
	}

	return 0;
}
