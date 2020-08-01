#include <bits/stdc++.h>

using namespace std;

// Complete the countingValleys function below.
int countingValleys(int n, string s)
{

    bool isMountain = true;
    bool isSeaLevel = true;
    int numberDown = 0;
    int numberUp = 0;
    int numberOfValley = 0;

    for (int i = 0; i < n; i++)
    {
        // check is now at sea level
        if (isSeaLevel)
        {
            if (s[i] == 'D')
            {
                isMountain = false;
                numberOfValley += 1;
                numberDown += 1;
                isSeaLevel = false;
            }
            else
            {
                numberUp += 1;
                isSeaLevel = false;
            }
        }
        else if (s[i] == 'D')
        {
            numberDown += 1;
        }
        else
        {
            numberUp += 1;
        }

        if (numberUp == numberDown)
        {
            numberUp = 0;
            numberDown = 0;
            isSeaLevel = true;
            if (!isMountain)
            {
                isMountain = true;
            }
        }
    }

    return numberOfValley;
}

int main()
{

    int n;
    cin >> n;
    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    string s;
    getline(cin, s);

    int result = countingValleys(n, s);

    cout << result << "\n";

    return 0;
}
