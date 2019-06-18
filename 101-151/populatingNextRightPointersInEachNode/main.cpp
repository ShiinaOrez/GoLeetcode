#include <iostream>
#include <cstdio>

using namespace std;

class Node {
public:
    int val;
    Node* left;
    Node* right;
    Node* next;

    Node() {}

    Node(int _val, Node* _left, Node* _right, Node* _next) {
        val = _val;
        left = _left;
        right = _right;
        next = _next;
    }
};

class Solution {
public:
    Node* deep[20];
    Node* connect(Node* root) {
        if (root != nullptr) {
            dfs(root, 0);
        }
        return root;
    }

    void dfs(Node* now, int d) {
        if (now->left != nullptr && now->right != nullptr) {
            now->left->next = now->right;
            now->right->next = nullptr;
            if (this->deep[d] != nullptr) {
                this->deep[d]->next = now->left;
            }
            this->deep[d] = now->right;
            dfs(now->left, d+1);
            dfs(now->right, d+1);
        }
        return;
    }
};

int main() {
    cout << endl;
}
