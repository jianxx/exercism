#include "pangram.h"

#include <cctype>
#include <unordered_set>
namespace pangram {
bool is_pangram(const string& word) {
    unordered_set<char> charSet;
    for (auto i = word.cbegin(); i != word.cend(); ++i) {
        char c = tolower(*i);
        if (isalpha(c)) {
            charSet.insert(c);
        }
    }
    return charSet.size() == 26;
}
}  // namespace pangram
