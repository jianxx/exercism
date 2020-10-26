#include "raindrops.h"

namespace raindrops {
std::string convert(int i) {
    std::string r("");
    if (i % 3 == 0) {
        r += "Pling";
    }
    if (i % 5 == 0) {
        r += "Plang";
    }
    if (i % 7 == 0) {
        r += "Plong";
    }
    if (r.compare("") == 0) {
        return std::to_string(i);
    } else {
        return r;
    }
}
}  // namespace raindrops
