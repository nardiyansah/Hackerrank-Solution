[Find the Point problem](https://www.hackerrank.com/challenges/find-point/problem)

My Solution
```
vector<int> findPoint(int px, int py, int qx, int qy)
{
  // determine the point r
  vector<int> r = {qx + (qx - px), qy + (qy - py)};

  return r;
}
```