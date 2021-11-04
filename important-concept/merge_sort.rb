# MergeSort(arr[], l,  r)
# If r > l
#      1. Find the middle point to divide the array into two halves:  
#              middle m = l+ (r-l)/2
#      2. Call mergeSort for first half:   
#              Call mergeSort(arr, l, m)
#      3. Call mergeSort for second half:
#              Call mergeSort(arr, m+1, r)
#      4. Merge the two halves sorted in step 2 and 3:
#              Call merge(arr, l, m, r)

def merge(arr, start_i, mid_i, end_i)
    sorted_arr = Array.new(end_i - start_i + 1, 0)
    i, j, k = start_i, mid_i, 0
    while i <= mid_i && j <= end_i
        if arr[i] < arr[j]
            sorted_arr[k] = arr[i]
            k += 1
            i += 1
        else
            sorted_arr[k] = arr[j]
            k += 1
            j += 1
        end
    end
    while i <= mid_i
        sorted_arr[k] = arr[i]
        k += 1
        i += 1
    end
    while j <= end_i
        sorted_arr[k] = arr[j]
        k += 1
        j += 1
    end
    sorted_arr
end

def merge_sort(arr, start_i, end_i)
    sorted_arr = []
    if start_i < end_i
        mid = (start_i + end_i) / 2
        merge_sort(arr, start_i, mid)
        merge_sort(arr, mid+1, end_i)
        sorted_arr = merge(arr, start_i, mid, end_i)
    end
    sorted_arr
end

arr = [64, 25, 12, 22, 11]
print arr
puts
sorted_arr = merge_sort(arr, 0, arr.length-1)
print sorted_arr
