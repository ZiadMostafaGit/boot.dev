#include<stdio.h>


void print_hash(int num){


for (int i=0;i<num;i++){
printf("# ");

}
printf("\n");





}

int main(){
int nums[]={5,1,4,1,1,1};

int size=sizeof(nums)/4;
for (int i=0;i<size;i++){
    print_hash(nums[i]);
}









}




