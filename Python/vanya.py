friend, h = map(int, input().split())
friend_h = list(map(int, input().split()))
sum = 0
for i in friend_h:
    if i <= h:
        sum += 1
    else:
        sum += 2
print(sum)
