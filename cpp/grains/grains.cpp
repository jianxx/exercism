#include "grains.h"

#include <limits>
#include <stdexcept>

namespace grains {
uint64_t square(int i) {
    if (i < 1 || i > 64) throw std::domain_error("Square must be between 1 and 64!");
    return uint64_t(1) << (i - 1);
}

uint64_t total() {
    return std::numeric_limits<uint64_t>::max();
}

}  // namespace grains
