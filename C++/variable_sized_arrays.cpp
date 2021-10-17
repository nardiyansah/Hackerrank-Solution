#include <cmath>
#include <cstdio>
#include <vector>
#include <iostream>
#include <algorithm>
using namespace std;


int main() {
    /* Enter your code here. Read input from STDIN. Print output to STDOUT */
  int n = 0;
  cin >> n;
  int q = 0;
  cin >> q;
  
  vector<vector<int>> vec;
  
  for (int i=0; i<n; i++) {
    vector<int> v1;
    int a;
    cin >> a;
    for (int j=0; j<a; j++) {
      int b;
      cin >> b;
      v1.push_back(b);
    }
    vec.push_back(v1);
  }
  for (int i=0; i<q; i++) {
    int j;
    cin >> j;
    int k;
    cin >> k;
    cout << vec[j][k] << endl;
  }
  return 0;
}
