#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <fcntl.h>
#include <unistd.h>

unsigned is_prime(unsigned n) {
	if (n <= 3) return n > 1;
	if (0 == (n % 2) || 0 == (n % 3)) return 0;

	unsigned i;
	for (i = 5; (i * i) <= n; i += 6)
		if (0 == (n % i) || 0 == (n % (i + 2))) return 0;

	return 1;
}

int main() {
	const char* file = "./fifoChannel";
	int fb = open(file, O_RDONLY);
	if (fb < 0) return -1;
	unsigned count = 0, total = 0, primes_count = 0;

	while (1) {
		int next;
		int i;

		ssize_t count = read(fb, &next, sizeof(int));
		if (0 == count) break;
		else if (count == sizeof(int)) {
			total++;
			if (is_prime(next)) primes_count++;
		}
	}

	close(fb);
	unlink(file);
	printf("Received ints: %u, primes: %u\n", total, primes_count);

	return 0;
}
