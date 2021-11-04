# iterate through index array and looking for the smallest value and swap
# time complexity: n^2
def selection_sort(arr)
    i = 0
    until i == arr.length - 1 do
        i_smallest_value = i
        for j in i..arr.length-1 do
            if arr[j] < arr[i_smallest_value]
                temp = arr[j]
                arr[j] = arr[i_smallest_value]
                arr[i_smallest_value] = temp
            end
        end
        i += 1
    end
    arr
end

arr = [64, 25, 12, 22, 11]
print arr
puts
sorted_arr = selection_sort(arr)
print sorted_arr