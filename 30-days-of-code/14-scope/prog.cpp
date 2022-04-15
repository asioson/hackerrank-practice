#include <cmath>
#include <cstdio>
#include <vector>
#include <iostream>
#include <algorithm>

using namespace std;

class Difference {
    private:
    vector<int> elements;
  
  	public:
  	int maximumDifference;

    Difference(vector<int> &a) {
        for (unsigned int i = 0; i < a.size(); i++) {
            elements.push_back(a[i]);
        }
    }

	void computeDifference() {
        maximumDifference = 0;
        if (elements.size() > 0) {
            int minVal, maxVal;
            minVal = maxVal = elements[0];
            for (long unsigned int i = 1; i < elements.size(); i++) {
                if (elements[i] > maxVal) 
                    maxVal = elements[i];
                else if (elements[i] < minVal)
                    minVal = elements[i];
            }
            maximumDifference = maxVal - minVal;
        }
    }

}; // End of Difference class

int main() {
    int N;
    cin >> N;
    
    vector<int> a;
    
    for (int i = 0; i < N; i++) {
        int e;
        cin >> e;
        
        a.push_back(e);
    }
    
    Difference d(a);
    
    d.computeDifference();
    
    cout << d.maximumDifference;
    
    return 0;
}
