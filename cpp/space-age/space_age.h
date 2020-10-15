#if !defined(SPACE_AGE_H)
#define SPACE_AGE_H

#include <functional>

namespace space_age {
using lamda = std::function<double()>;

class space_age {
    unsigned long secs;
    lamda func(double factor);

   public:
    space_age(unsigned long seconds);
    unsigned long seconds() const;
    double on_earth() const;
    lamda on_neptune;
    lamda on_uranus;
    lamda on_venus;
    lamda on_mercury;
    lamda on_mars;
    lamda on_jupiter;
    lamda on_saturn;
};
}  // namespace space_age

#endif  // SPACE_AGE_H