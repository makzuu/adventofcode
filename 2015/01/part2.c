#include <stdio.h>

int main() {
    int c;
    int floor = 0;
    int position = 1;
    while ((c = getchar()) != EOF) {
        switch (c) {
            case '(':
                floor++;
                break;
            case ')':
                floor--;
                break;
        }
        if (floor == -1) {
            printf("%d\n", position);
            break;
        }
        position++;
    }
}
