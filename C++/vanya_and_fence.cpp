#include <iostream>

using namespace std;

int main() {
	int n_friend;
	cin >> n_friend;
	int h;
	cin >> h;
	int f_h[n_friend];
	for(int i=0; i<n_friend; i++){
		cin >> f_h[i];
		cout << f_h[i];
	}
}
