#include <stdio.h>
#include <stdlib.h>
#include <CoreGraphics/libcg.h>

int main() {
	char *r = CGMainDisplay();
	printf("\nGPU: %s\n", r);
	free(r);
}
