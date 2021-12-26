class Solution:
    def displayTable(self, orders: List[List[str]]) -> List[List[str]]:
        table = {}
        foods = []
        answer = []
        for order in orders:
            table_number = order[1]
            food = order[2]
            if food not in foods:
                foods.append(food)
            if table_number in table.keys():
                if food in table[table_number].keys():
                    table[table_number][food] += 1
                else:
                    table[table_number][food] = 1
            else:
                table[table_number] = {food: 1}
        order_table = list(map(int, table.keys()))
        order_table.sort()
        foods.sort()
        first_row = ["Table"]
        for f in foods:
            first_row.append(f)
        answer.append(first_row)
        for n in order_table:
            t = str(n)
            new_table = []
            new_table.append(t)
            for i in range(1,len(first_row)):
                if first_row[i] not in table[t].keys():
                    new_table.append("0")
                else:
                    new_table.append(str(table[t][first_row[i]]))
            answer.append(new_table)
        return answer
