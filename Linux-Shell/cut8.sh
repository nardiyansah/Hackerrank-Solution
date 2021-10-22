while read line
do
  echo "$line" | cut -d " " -f 1,2,3
done
