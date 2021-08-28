#!/bin/ruby

require 'json'
require 'stringio'

#
# Complete the 'rotLeft' function below.
#
# The function is expected to return an INTEGER_ARRAY.
# The function accepts following parameters:
#  1. INTEGER_ARRAY a
#  2. INTEGER d
#

def rotLeft(a, d)
    # Write your code here
    i = 0
    while i < d
        first_element = a[0]
        temp_arr = a[1..-1]
        temp_arr.append(first_element)
        a = temp_arr
        i = i + 1
    end
    a
end

fptr = File.open(ENV['OUTPUT_PATH'], 'w')

first_multiple_input = gets.rstrip.split

n = first_multiple_input[0].to_i

d = first_multiple_input[1].to_i

a = gets.rstrip.split.map(&:to_i)

result = rotLeft a, d

fptr.write result.join " "
fptr.write "\n"

fptr.close()
