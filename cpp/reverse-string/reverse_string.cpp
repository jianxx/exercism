#include "reverse_string.h"

namespace reverse_string {
using std::reverse;
string reverse_string(string input) {
    reverse(input.begin(), input.end());
    return input;
}
}  // namespace reverse_string
