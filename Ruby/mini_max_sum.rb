def miniMaxSum(arr)
    # sort first
    arr.sort!
    printf("%d %d", arr[0..3].sum, arr[-4, 4].sum)
end