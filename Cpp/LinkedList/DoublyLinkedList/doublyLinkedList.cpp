#include <iostream>

using namespace std;

struct node {
    int data;
    struct node* next;
    struct node* prev;
} *head;

node* createNode(int val){
	struct node *node, *n;

	node = new(struct node);

	node->data = val;
	node->next = NULL;
	node->prev = NULL;

	head = node;

	return node;
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

void insertLast()
{
    struct node *n, *h;
    int val;
    cout << "value: ";
    cin >> val;

	if (head == NULL){
		n = createNode(val);
		return;
	}else {
		n = new(struct node);
		h = head;
  	 	while (h->next != NULL)
  	  	{
 	       h = h->next;
	    }

  		h->next = n;
   		n->next = NULL;
		n->data = val;
		n->prev = h;
	}    
}

void insertInPosition(){
	struct node  *n, *p;

	int val, pos;
	cout << "Value: ";
    cin >> val;
	cout << "Position: ";
    cin >> pos;

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

void deleteInPosition()
{
    struct node *s, *prev;
    int pos, i, count = 0;

    if (head == NULL)
        return;

    cout << "position: ";
    cin >> pos;

    s = head;
    if (pos == 1)
        head = s->next;
    else
    {
        while (s != NULL)
        {
            s = s->next;
            count++;
        }

        if (pos >= 2 && pos <= count)
        {
            s = head;
            for (i = 1; i <= pos - 1; i++)
            {
                prev = s;
                s = s->next;
            }
            prev->next = s->next;
        }
        else
            cout << "out of range" << endl;

        free(s);
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
            insertLast();
            cout << endl;
            break;
        case 3:
            insertInPosition();
            cout << endl;
            break;
        case 4:
            deleteInPosition();
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