#include "pangram.h"

#include <set>
namespace pangram {
bool is_pangram(const string word) {
    set<char> charSet;
    for (auto i = word.cbegin(); i != word.cend(); ++i) {
        char c = tolower(*i);
        if (c >= 'a' && c <= 'z') {
            charSet.insert(c);
        }
    }
    return charSet.size() == 26;
}
}  // namespace pangram
