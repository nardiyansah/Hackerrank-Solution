[Sock Merchant Problem](https://www.hackerrank.com/challenges/sock-merchant/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=warmup)

My Solution
```
// Complete the sockMerchant function below.
int sockMerchant(int n, vector<int> ar) {
    // number of pairs
    int pair=0;
    // sort the vector
    sort(ar.begin(), ar.end());

    for(vector<int>::iterator i=ar.begin(); i != ar.end();){
        vector<int>::iterator j = i+1;
        
        if(*i == *j){
            pair = pair + 1;
            i = i + 2;
        }else{
            i = i + 1;
        }
    }

return pair;
}
```
