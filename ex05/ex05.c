//
// 5) Escreva um programa que inverta os caracteres de um string.

// IMPORTANTE:
// a) Essa string pode ser informada através de qualquer entrada de sua preferência ou pode ser previamente definida no código;
// b) Evite usar funções prontas, como, por exemplo, reverse;

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

static void reverse(char *s)
{
    int i = 0;
    int j = strlen(s) - 1;
    char c = 0;

    while (i < j)
    {
        c = s[i];
        s[i++] = s[j];
        s[j--] = c;
    }
}

int main(int argc, char **argv)
{
    if (argc != 2)
        return 1;
    reverse(argv[1]);
    printf("%s\n", argv[1]);
    return 0;
}