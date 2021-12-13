#include <bits/stdc++.h>

using namespace std;

string ltrim(const string &);
string rtrim(const string &);
vector<string> split(const string &);

vector<int> countingSort(vector<int> arr)
{
    vector<int> frequency_arr(100, 0);
    int i_value = 0;

    for (int i = 0; i < arr.size(); i++)
    {
        i_value = arr[i];
        frequency_arr[i_value] += 1;
    }

    return frequency_arr;
}