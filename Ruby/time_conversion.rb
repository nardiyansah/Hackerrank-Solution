def timeConversion(s)
    type_time = s[-2, 2]
    conversion_time = ""
    if type_time == "AM"
        if s[0..1] == "12"
            conversion_time = "00" + s[2..7]
        else
            conversion_time = s[0..7]
        end
    else
        if s[0..1] == "12"
            conversion_time = s[0..7]
        else
            hour = s[0..1]
            hour = hour.to_i
            hour = hour + 12
            hour = hour.to_s
            conversion_time = hour + s[2..7]
        end
    end
    return conversion_time
end