#include <iostream>
#include <cstdlib>
#include <cstdio>
#include <climits>

using namespace std;

struct node
{
    int data;
    struct node* next;
} *head;

node* createNode(int value)
{
    struct node* node;
    node = new (struct node);

    if (node == NULL)
    {
        cout << "not allocated" << endl;
        return 0;
    }
    else
    {
        node->data = value;
        node->next = NULL;
        return node;
    }
}

void insertBegin()
{
    struct node *n, *prev;
    int val;
    cout << "value: ";
    cin >> val;

    n = createNode(val);

    if (head == NULL)
    {
        head = n;
        head->next = NULL;
    }
    else
    {
        prev = head;
        head = n;
        head->next = prev;
    }
}

int main(){
     int num;

    // Start with empty list
    head = NULL;

    while (1)
    {
        cout << "1. Insert in begin" << endl;
        cout << "2. Exit" << endl;
        cout << "\n Enter a number between 1 and 2: " << endl;

        cin >> num;

        switch (num)
        {
        case 1:
            insertBegin();
            cout << endl;
            break;

        default:
            cout << "Please enter a valid number between 1 to 2";
            break;
        }
    }

    return 0;
}