#include <bits/stdc++.h>

using namespace std;

string gridChallenge(vector<string> grid) {
    if (grid.size() == 1) {
        return "YES";
    }
    for (int i = 0; i<grid.size(); i++) {
        sort(grid[i].begin(), grid[i].end());
    }
    for (int i = 0; i<grid.size(); i++) {
        for (int j=0; j<grid.size()-2; j++) {
            if (grid[i][j] > grid[i][j+1]) {
                cout << grid[i][j+1] << endl; 
                return "NO";
            }
            if (grid[j][i] > grid[j+1][i]) {
                return "NO";
            }
        }
    }
    return "YES";
}