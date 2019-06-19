#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>
#include <time.h>
#include <stdlib.h>
#include <stdio.h>

#define MaxLoops     12000
#define ChunkSize    16
#define IntsPerChunk 4
#define MaxZs        250

int main() {
	const char* pipeName = "./fifoChannel";
	mkfifo(pipeName, 0666);
	int fb = open(pipeName, O_CREAT | O_WRONLY);
	if (fb < 0) return -1;

	int i;
	for (i = 0; i < MaxLoops; i++) {
		int j;
		for (j = 0; j < ChunkSize; j++) {
			int k;
			int chunk[IntsPerChunk];
			for (k = 0; k < IntsPerChunk; k++)
				chunk[k] = rand();
			write(fb, chunk, sizeof(chunk));
		}
		usleep((rand() % MaxZs) + 1);
	}

	close(fb);
	unlink(pipeName);
	printf("%i ints sent to the pipe.\n", MaxLoops * ChunkSize * IntsPerChunk);

	return 0;
}
