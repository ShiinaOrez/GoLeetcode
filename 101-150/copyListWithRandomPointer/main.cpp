#include<cstdio>
#include<iostream>

using namespace std;

class Node {
public:
    int val;
    Node* next;
    Node* random;

    Node() {}

    Node(int _val, Node* _next, Node* _random) {
        val = _val;
        next = _next;
        random = _random;
    }
};

class Solution {
public:
    Node* copyRandomList(Node* head) {
        if (head == nullptr) {
            return head;
        }
        Node* node = head;
        while(node != nullptr) {
            Node* clone = new Node(node->val, node->next, nullptr);
            Node* temp = node->next;
            node->next = clone;
            node = temp;
        }
        node = head;
        while(node != nullptr) {
            node->next->random = node->random == nullptr ? nullptr : node->random->next;
            node = node->next->next;
        }
        node = head;
        Node* cloneHead = head->next;
        while(node->next != nullptr) {
            Node* temp = node->next;
            node->next = node->next->next;
            node = temp;
        }
        return cloneHead;
    }
};

int main() {
    printf("\n");
}
