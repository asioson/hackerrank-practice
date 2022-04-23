#include <iostream>
#include <cstddef>
#include <queue>
#include <string>
#include <cstdlib>

using namespace std;	
class Node{
    public:
        int data;
        Node *left,*right;
        Node(int d){
            data=d;
            left=right=NULL;
        }
};
class Solution{
    public:
  		Node* insert(Node* root, int data){
            if(root==NULL){
                return new Node(data);
            }
            else{
                Node* cur;
                if(data<=root->data){
                    cur=insert(root->left,data);
                    root->left=cur;
                }
                else{
                   cur=insert(root->right,data);
                   root->right=cur;
                 }           
           return root;
           }
        }

    void printLevel(Node *root, int level, int target) {
        if (root == NULL) return;
        if (level > target) return;
        if (level == target) 
            cout << root->data << " ";
        printLevel(root->left, level+1, target);
        printLevel(root->right, level+1, target);
    }
    
    int getMaxLevel(Node *root) {
        if (root == NULL) return 0;
        return 1 + max(getMaxLevel(root->left), getMaxLevel(root->right));
    }

	void levelOrder(Node * root){
        if (root == NULL) return;
        int n = getMaxLevel(root); 
        for (int i = 0; i < n; i++) 
            printLevel(root, 0, i);
	}

};//End of Solution
int main(){
    Solution myTree;
    Node* root=NULL;
    int T,data;
    cin>>T;
    while(T-->0){
        cin>>data;
        root= myTree.insert(root,data);
    }
    myTree.levelOrder(root);
    return 0;
}
