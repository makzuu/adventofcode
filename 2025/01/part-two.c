#include <stdio.h>
#include <stdlib.h>

int rotate(int dial, int rotation) {
    return ((dial + rotation) % 100 + 100) % 100;
}

int main() {
    int dial = 50;
    int password = 0;

    int c;
    char buff[50];
    int i = 0;
    while ((c = getchar()) != EOF) {
	if (c != '\n') {
	    buff[i++] = c;
	} else {
	    buff[i] = '\0';
	    i = 0;

	    char direction = buff[0];
	    int distance = atoi(buff + 1);

	    for (; distance != 0; distance--) {
		if (direction == 'L') {
		    dial = rotate(dial, -1);
		    if (dial == -1) {
			dial = 99;
		    }
		} else {
		    dial = rotate(dial, 1);
		    if (dial == 100) {
			dial = 0;
		    }
		}
		if (dial == 0) {
		    password++;
		}
	    }
	}
    }
    printf("%d\n", password);
}
