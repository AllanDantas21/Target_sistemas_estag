// 2) Dado a sequência de Fibonacci, onde se inicia por 0 e 1 e o próximo valor sempre será a soma dos 2 valores anteriores 
//  (exemplo: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34...), escreva um programa na linguagem que desejar onde, informado um número, 
// ele calcule a sequência de Fibonacci e retorne uma mensagem avisando se o número informado pertence ou não a sequência.

// IMPORTANTE: Esse número pode ser informado através de qualquer entrada de sua preferência ou pode ser previamente definido no código;

#include <stdio.h>
#include <stdlib.h>

static int is_fibonacci(int n, long a, long b)
{
    if (n == a || n == b)
        return 1;
    if (b > n)
        return 0;
    return is_fibonacci(n, b, a + b);
}

static void fibonacci(long n)
{
    if (is_fibonacci(n, 0, 1))
        printf("O número %ld pertence à sequência de Fibonacci\n", n);
    else
        printf("O número %ld não pertence à sequência de Fibonacci\n", n);
}

int main(int argc, char **argv)
{
    if (argc == 2)
        fibonacci(atoi(argv[1]));
    return 0;
}