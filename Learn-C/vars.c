#include<stdio.h>
#include<string.h>

#define true 1
#define false 0








int main(){

char *name="Ziad_mostafa_el7said";


int size=strlen(name);


int flag=false;


for (int i=0;i<size;i++){

if (name[i]=='7'){
flag=true;



}


}


if (flag==true){


printf("contain 7");



}else{

printf("dose not contain 7");




}


return 0;


}



