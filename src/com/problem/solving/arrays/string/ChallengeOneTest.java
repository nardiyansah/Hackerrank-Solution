package com.problem.solving.arrays.string;

import static org.junit.jupiter.api.Assertions.*;

class ChallengeOneTest {

    @org.junit.jupiter.api.Test
    void uniqueChars() {
        ChallengeOne challengeOne = new ChallengeOne();

        assertFalse(challengeOne.uniqueChars());
        assertTrue(challengeOne.uniqueChars(""));
        assertFalse(challengeOne.uniqueChars("foo"));
        assertTrue(challengeOne.uniqueChars("bar"));
    }
}