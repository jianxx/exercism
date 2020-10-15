#if !defined(NUCLEOTIDE_COUNT_H)
#define NUCLEOTIDE_COUNT_H
#include <map>
#include <stdexcept>
#include <string>

namespace nucleotide_count {
class counter {
    std::map<char, int> _counts;

   public:
    counter(const std::string& s);
    std::map<char, int> nucleotide_counts() const;
    int count(char c) const;
};
}  // namespace nucleotide_count

#endif  // NUCLEOTIDE_COUNT_H