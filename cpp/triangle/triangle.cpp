#include "triangle.h"

#include <algorithm>
#include <cmath>
#include <iomanip>
#include <iostream>
#include <limits>
#include <stdexcept>
#include <type_traits>

namespace triangle {
template <class T>
typename std::enable_if<!std::numeric_limits<T>::is_integer, bool>::type
almost_equal(T x, T y, int ulp) {
    return std::fabs(x - y) <= std::numeric_limits<T>::epsilon() * std::fabs(x + y) * ulp || std::fabs(x - y) < std::numeric_limits<T>::min();
}

flavor kind(double a, double b, double c) {
    if (b > c) std::swap(b, c);
    if (a > c) std::swap(a, c);
    if (a > b) std::swap(a, b);
    if (a <= 0) throw std::domain_error("All sides have to be of length > 0!");
    if (c > a + b) throw std::domain_error("The sum of the lengths of any two sides must be greater than or equal to the length of the third side!");

    if (almost_equal(a, b, 2) && almost_equal(b, c, 2))
        return flavor::equilateral;
    else if (almost_equal(a, b, 2) || almost_equal(b, c, 2))
        return flavor::isosceles;
    else
        return flavor::scalene;
}
}  // namespace triangle
