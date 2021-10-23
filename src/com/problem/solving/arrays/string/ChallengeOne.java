package com.problem.solving.arrays.string;

import java.util.HashMap;

public class ChallengeOne {

    public boolean uniqueChars(String str) {
        if (str == null) return false;
        HashMap<Character, Integer> character = new HashMap<Character, Integer>();
        for (int i=0; i<str.length(); i++) {
            char s = str.charAt(i);
            if (character.containsKey(s)) return false;
            character.put(s, 1);
        }
        return true;
    }

    public boolean uniqueChars() {
        return uniqueChars(null);
    }
}
