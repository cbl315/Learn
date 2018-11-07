#include<stdio.h>

/* 
 getchar和putchar的使用
*/

int main(){
    int c;
    printf("\nEOF:%c\n", EOF);
    while((c = getchar()) != EOF){
        putchar(c); 
    }
    return 0;
}