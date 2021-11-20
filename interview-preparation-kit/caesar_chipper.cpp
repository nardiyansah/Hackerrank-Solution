#include <bits/stdc++.h>

using namespace std;

string caesarCipher(string s, int k) {
    string encrypted_s = "";
    k = k % 26;
    for(auto c : s) {
        if(int(c) >= int('A') && int(c) <= int('Z')) {
            int encrypted_ascii_c = int(c) + k;
            if (encrypted_ascii_c > int('Z')) {
                if ((encrypted_ascii_c % int('A')) == 0) {
                    encrypted_ascii_c = int('A');
                } else {
                    int remainder = (encrypted_ascii_c % int('A'));
                    if (remainder > 26) {
                        remainder = remainder % 26;
                    }
                    if (remainder == 26) {
                      remainder = 0;
                    }
                    encrypted_ascii_c = remainder + int('A');
                }
            }
            char encrypted_c = encrypted_ascii_c;
            encrypted_s += encrypted_c;
        } else if (int(c) >= int('a') && int(c) <= int('z')) {
            int encrypted_ascii_c = int(c) + k;
            if (encrypted_ascii_c > int('z')) {
                if ((encrypted_ascii_c % int('a')) == 0) {
                    encrypted_ascii_c = int('a');
                } else {
                    int remainder = (encrypted_ascii_c % int('a'));
                    if (remainder > 26) {
                        remainder = remainder % 26;
                    }
                    if (remainder == 26) {
                      remainder = 0;
                    }
                    encrypted_ascii_c = remainder + int('a');
                }
            }
            char encrypted_c = encrypted_ascii_c;
            encrypted_s += encrypted_c;
        } else {
            encrypted_s += c;
        }
    }
    return encrypted_s;
}
