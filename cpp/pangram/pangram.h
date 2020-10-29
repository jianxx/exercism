#if !defined(PANGRAM_H)
#define PANGRAM_H
#include <string>
namespace pangram {
using namespace std;
bool is_pangram(const string& word);
}  // namespace pangram

#endif  // PANGRAM_H