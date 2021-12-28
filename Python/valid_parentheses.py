class Solution:
    def isValid(self, s: str) -> bool:
        if len(s) % 2 == 1:
            return False
        stackArr = []
        try:
            for bracket in s:
                if (
                    (bracket == ")" and stackArr[-1] == "(")
                    or (bracket == "}" and stackArr[-1] == "{")
                    or (bracket == "]" and stackArr[-1] == "[")
                ):
                    stackArr.pop()
                else:
                    stackArr.append(bracket)
        except:
            return False
        if len(stackArr) == 0:
            return True
        else:
            return False
