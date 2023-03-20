#include <iostream>

using namespace std;

struct node {
    int data;
    struct node* next;
    struct node* prev;
} *head;

void listMenu(){
    cout << "1. Insert in begin" << endl;
    cout << "2. Insert in last" << endl;
    cout << "3. Insert in position" << endl;
    cout << "4. Delete in position" << endl;
    cout << "5. Dispaly linkedlist" << endl;
    cout << "6. Exit" << endl;
    cout << "\n Enter a number between 1 and 6: " << endl;
}

int main(){
    listMenu();
    return 0;
}