// cc part1.c && ./a.out < input

#include <stdio.h>

int main() {
    int c;
    int floor = 0;
    while ((c = getchar()) != EOF) {
        switch (c) {
            case '(':
                floor++;
                break;
            case ')':
                floor--;
                break;
        }
    }
    printf("%d\n", floor);
}
