void whatFlavors(vector<int> cost, int money) {
    // key : element value
    // value : element index
    std::unordered_map<int, int> hashCost;
    for (int i=0; i<cost.size(); i++) {
        int remain = money - cost[i];
        if (hashCost.find(remain)!=hashCost.end()) {
            std::unordered_map<int, int>::iterator it = hashCost.find(remain);
            cout << it->second+1 << " " << i+1 << endl;
            break;
        }
        hashCost[cost[i]] = i;
    }
}