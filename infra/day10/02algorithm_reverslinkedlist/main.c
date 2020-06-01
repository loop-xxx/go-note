#include <stdio.h>

typedef struct _linkednode{
    unsigned int id ;
    struct _linkednode *next ;
} linkednode ;


void revers_linkedlist(linkednode *head){
    linkednode *prev, *curr, *next ;
    
    prev = NULL ;
    next = head ;
    
    while((curr = next)){
        next = curr->next ;
        curr->next = prev ;
        prev = curr ;
    }

}
void show_linkedlist(linkednode *head){
    linkednode *curr = head ;
    
    while(curr){
        printf("%u\n", curr->id) ;
        curr = curr->next ;
    }

}


int main(int argc, char **argv){
    linkednode l1 = {0, NULL} ;
    linkednode l2 = {1, NULL} ;
    linkednode l3 = {2, NULL} ;
    linkednode l4 = {3, NULL} ;
    linkednode l5 = {4, NULL} ;
    linkednode l6 = {5, NULL} ;

    l1.next = &l2 ;
    l2.next = &l3 ;
    l3.next = &l4 ;
    l4.next = &l5 ;
    l5.next = &l6 ;
    show_linkedlist(&l1) ;
    revers_linkedlist(&l1) ;
    show_linkedlist(&l6) ;

    return 0 ;
}