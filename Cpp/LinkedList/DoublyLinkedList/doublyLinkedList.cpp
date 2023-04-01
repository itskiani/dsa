#include <iostream>

using namespace std;

struct node {
    int data;
    struct node* next;
    struct node* prev;
} *head;

void createNode(int val){
	struct node *node, *n;

	node = new(struct node);
	node->data = val;
	node->next = NULL;

	if (head == NULL){
		node->prev = NULL;
		head = node;
	} else {
		return;
	}
}

void insertBegin(){
	int val;
	cout << "Value: ";
	cin >> val;

	if (head == NULL){
		createNode(val);
		return;
	}

	struct node *t;
	t = new(struct node);
	t->data = val;
	t->prev = NULL;
	t->next = head;

	head->prev = t;
	head = t;
}

void insertInPosition(int val, int pos){
	struct node  *n, *p;

	if (head == NULL)
		return;
	
	p = head;
	for (int i=0; i<pos-1; i++){
		p = p->next;
		if(p == NULL){
			cout << "Sorry, the position you entered is wrong.";
		}
	}

	n = new(struct node);
	n->data = val;
	if(p->next == NULL){
		p->next = n;
		n->next = NULL;
		n->prev = p;
	}else{
		n->next = p->next;
		p->next->prev = n;
		p->next = n;
		n->prev = p;
	}
}

void display()
{
    struct node *s;
    if (head == NULL)
    {
        cout << "Empty" << endl;
        return;
    }

    s = head;
    cout << "Elements: " << endl;
    while (s != NULL)
    {
        cout << s->data << "<->";
        s = s->next;
    }
    cout << "NULL" << endl;
}

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
	int num;

    // Start with empty list
    head = NULL;

    while (1)
    {
        listMenu();
        cin >> num;

        switch (num)
        {
        case 1:
            insertBegin();
            cout << endl;
            break;
        case 2:
            // insertLast();
            cout << endl;
            break;
        case 3:
            // insertInPosition();
            cout << endl;
            break;
        case 4:
            // deleteInPosition();
            cout << endl;
            break;
        case 5:
            display();
            cout << endl;
            break;
        case 6:
            exit(1);
            cout << endl;
            break;    

        default:
            cout << "Please enter a valid number between 1 to 6";
            break;
        }
    }

    return 0;

}