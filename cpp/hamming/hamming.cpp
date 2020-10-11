#include "hamming.h"

namespace hamming {
int compute(const string a, const string b) {
    int count = 0;
    for (auto i = a.cbegin(), j = b.cbegin(); i != a.cend() && j != b.cend();) {
        if (*i != *j) {
            count++;
        }
        ++i;
        ++j;
    }
    return count;
}
}  // namespace hamming
