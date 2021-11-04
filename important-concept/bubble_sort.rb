# using two loop
# if element i > element i+1 then swap

def bubble_sort(arr)
    for i in 0..arr.length - 1 do
        # last i element already in place
        for j in 0..arr.length - (i + 2) do
            if arr[j] > arr[j+1]
                temp = arr[j]
                arr[j] = arr[j+1]
                arr[j+1] = temp
            end
        end
    end
    arr
end

arr = [64, 25, 12, 22, 11]
print arr
puts
sorted_arr = bubble_sort(arr)
print sorted_arr