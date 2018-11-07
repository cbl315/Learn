#include <stdio.h>

/*
 变量的使用
*/


int main(){
    float fahr, celsius;
    int lower, upper, step;
    lower = 0;  // 温度表的下限
    upper = 300;  // 温度表的上限
    step = 20;  // 步长
    printf("fahr celsius\n");
    fahr = lower;
    while (fahr <= upper) {
        celsius = (5.0/9.0) * (fahr - 32.0);
        printf("%3.0f %6.1f\n", fahr, celsius);
        fahr = fahr + step;
    }

    return 0;
}