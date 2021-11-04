#include <bits/stdc++.h>

using namespace std;

int lonelyinteger(vector<int> a) {
    unordered_map<int, int> valueMap;
    int sizeA = a.size();
    int uniq = 0;
    for (int i = 0; i < sizeA; i++)
    {
        if (valueMap.find(a[i]) == valueMap.end())
        {
            valueMap[a[i]] = 1;
        }
        else
        {
            valueMap[a[i]] += 1;
        }
    }
    for (auto x : valueMap) {
        if (x.second == 1) {
            uniq = x.first;
        }
    }
    return uniq;
}