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

void insertLast()
{
    struct node *n, *h;
    int val;
    cout << "value: ";
    cin >> val;

    n = createNode(val);
    h = head;

    while (h->next != NULL)
    {
        h = h->next;
    }

    n->next = NULL;
    h->next = n;
}

void insertInPosition()
{
    struct node *n, *s, *p;
    int val, pos, count = 0;
    cout << "data: ";
    cin >> val;

    n = createNode(val);
    cout << "position: ";
    cin >> pos;

    int i;
    s = head;

    while (s != NULL)
    {
        s = s->next;
        count++;
    }

    if (pos == 1)
    {
        if (head == NULL)
        {
            head = n;
            head->next = NULL;
        }
        else
        {
            p = head;
            head = n;
            head->next = p;
        }
    }
    else if (pos > 1 && pos <= count)
    {
        s = head;
        for (i = 1; i < pos; i++)
        {
            p = s;
            s = s->next;
        }
        p->next = n;
        n->next = s;
    }
    else
    {
        cout << "Position out of range." << endl;
    }
}

int main(){
     int num;

    // Start with empty list
    head = NULL;

    while (1)
    {
        cout << "1. Insert in begin" << endl;
        cout << "2. Insert in last" << endl;
        cout << "3. Insert in position" << endl;
        cout << "4. Exit" << endl;
        cout << "\n Enter a number between 1 and 4: " << endl;

        cin >> num;

        switch (num)
        {
        case 1:
            insertBegin();
            cout << endl;
            break;
        case 2:
            insertLast();
            cout << endl;
            break;
        case 3:
            insertInPosition();
            cout << endl;
            break;
        case 4:
            exit(1);
            cout << endl;
            break;    

        default:
            cout << "Please enter a valid number between 1 to 4";
            break;
        }
    }

    return 0;
}