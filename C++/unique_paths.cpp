class Solution {
    unordered_map<string, int> grid;
public:
    int uniquePaths(int m, int n) {
        if(m == 1 && n == 1) return 1;
        if(m == 0 || n == 0) return 0;
        string key = to_string(m) + "-" + to_string(n);
        if(grid.find(key) != grid.end()) return grid[key];
        return grid[key] = uniquePaths(m-1, n) + uniquePaths(m, n-1);
    }
};