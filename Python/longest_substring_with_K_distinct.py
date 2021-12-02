def longest_substring(complete_string, K):
    max_substring = 0
    substring = ""
    num_distinct = 0
    dict_char_substring = {}

    for character in complete_string:
        substring += character

        if character not in dict_char_substring.keys():
            num_distinct += 1
            dict_char_substring[character] = 1
        else:
            dict_char_substring[character] += 1

        while num_distinct > K:
            first_char = substring[0]
            substring = substring[1:]
            dict_char_substring[first_char] -= 1

            if dict_char_substring[first_char] == 0:
                dict_char_substring.pop(first_char)
                num_distinct -= 1

        if len(substring) > max_substring:
            max_substring = len(substring)

    return max_substring


max_length_substring = longest_substring("cbbebi", 3)
print(max_length_substring)
