def plusMinus(arr)
    # Write your code here
    length = arr.length
    positive = 0
    negative = 0
    zero = 0
    
    for i in arr do
        if i > 0 then positive += 1 end
        if i < 0 then negative += 1 end
        if i == 0 then zero += 1 end
    end
    
    puts (positive.to_f/length).round(6)
    puts (negative.to_f/length).round(6)
    puts (zero.to_f/length).round(6)
end