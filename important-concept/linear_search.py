def linear_search(list, target):
    for i in range(0, len(list)):
        if target == list[i]:
            return i
    return None
