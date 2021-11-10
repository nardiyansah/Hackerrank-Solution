#include <bits/stdc++.h>

using namespace std;

int diagonalDifference(vector<vector<int>> arr) {
    int n = arr[0].size();
    int upperLeftToBottomRight = 0;
    int upperRightToBottomLeft = 0;
    for (int i = 0; i < n; i++)
    {
        upperLeftToBottomRight += arr[i][i];
        upperRightToBottomLeft += arr[i][n-1-i];
    }
    return abs(upperRightToBottomLeft - upperLeftToBottomRight);
}

