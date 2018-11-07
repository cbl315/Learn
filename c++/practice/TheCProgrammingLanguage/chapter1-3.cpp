#include <stdio.h>

/*
 for循环的使用
*/


int main(){
    int fahr;
    printf("fahr\tcelsius\n");
    for(fahr=300;fahr >= 0;fahr -= 20){
        printf("%3d %6.2f\n", fahr, (5.0/9.0)*(fahr-32));
    }

    return 0;
}