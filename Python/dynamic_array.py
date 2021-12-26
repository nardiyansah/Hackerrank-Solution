# # Complete the 'dynamicArray' function below.
# #
# # The function is expected to return an INTEGER_ARRAY.
# # The function accepts following parameters:
# #  1. INTEGER n
# #  2. 2D_INTEGER_ARRAY queries
# #


def dynamicArray(n, queries):
    # Write your code here
    dyn_arr = []
    for i in range(n):
        dyn_arr.append([])
    last_answer = 0
    answer_arr = []
    for q in queries:
        idx = (q[1] ^ last_answer) % n
        if q[0] == 1:
            dyn_arr[idx].append(q[2])
        else:
            last_answer = dyn_arr[idx][q[2] % len(dyn_arr[idx])]
            answer_arr.append(last_answer)
    return answer_arr


n = 2
arr_query = [[1, 0, 5], [1, 1, 7], [1, 0, 3], [2, 1, 0], [2, 1, 1]]

print(dynamicArray(n, arr_query))
