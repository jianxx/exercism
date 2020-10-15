#include "nucleotide_count.h"

namespace nucleotide_count {
counter::counter(const std::string& s) : _counts(std::map<char, int>({{'A', 0},
                                                                      {'T', 0},
                                                                      {'C', 0},
                                                                      {'G', 0}})) {
    for (char c : s)
        ++_counts[c];
}
std::map<char, int> counter::nucleotide_counts() const {
    return _counts;
}

int counter::count(char c) const {
    if (_counts.find(c) == _counts.end())
        throw std::invalid_argument(std::string("invalid nucleotide: ") + c);

    return _counts.at(c);
}
}  // namespace nucleotide_count
