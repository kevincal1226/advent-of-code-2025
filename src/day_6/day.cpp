#include <algorithm>
#include <cassert>
#include <fstream>
#include <iostream>
#include <numeric>
#include <print>
#include <ranges>
#include <sstream>
#include <string>
#include <vector>

namespace {
struct Column {
    std::vector<uint64_t> nums;

    char operation {};
};

auto part_1(const char* filename) -> uint64_t {
    std::vector<Column> input;
    std::ifstream ifs { filename };

    std::string line;
    while (std::getline(ifs, line)) {
        std::stringstream ss { line };
        std::string val;
        size_t sz {};
        while (ss >> val) {
            if (sz >= input.size()) {
                input.emplace_back();
            }
            if (val.contains('+') || val.contains('*')) {
                input[sz++].operation = val[0];
                continue;
            } else {
                input[sz++].nums.emplace_back(std::stoull(val));
            }
        }
    }

    std::ranges::for_each(input, [](auto& a) { std::println("{} {}", a.operation, a.nums); });

    return std::accumulate(input.begin(), input.end(), 0ULL, [](uint64_t acc, Column& c) {
        return acc
             + std::accumulate(c.nums.begin(), c.nums.end(), c.operation == '*' ? 1ULL : 0ULL,
                               [&](uint64_t sub_acc, uint64_t n) {
                                   switch (c.operation) {
                                   case '*':
                                       return sub_acc * n;
                                   case '+':
                                       return sub_acc + n;
                                   default:
                                       assert(false);
                                   }
                               });
    });
}

auto part_2(const char* filename) -> uint64_t {
    std::vector<Column> input;

    std::ifstream ifs { filename };

    std::string line;
    while (std::getline(ifs, line)) {
        std::stringstream ss { line };
        std::string val;
        size_t sz {};
        while (ss >> val) {
            if (sz >= input.size()) {
                input.emplace_back();
            }
            if (val.contains('+') || val.contains('*')) {
                input[sz].operation = val[0];
            } else {
                input[sz].nums.emplace_back(std::stoull(val));
            }
            ++sz;
        }
    }

    std::ranges::for_each(input, [](auto& a) { std::println("{} {}", a.operation, a.nums); });

    return std::accumulate(input.begin(), input.end(), 0ULL, [](uint64_t acc, Column& c) {
        return acc
             + std::accumulate(c.nums.begin(), c.nums.end(), c.operation == '*' ? 1ULL : 0ULL,
                               [&](uint64_t sub_acc, uint64_t n) {
                                   switch (c.operation) {
                                   case '*':
                                       return sub_acc * n;
                                   case '+':
                                       return sub_acc + n;
                                   default:
                                       assert(false);
                                   }
                               });
    });
}

}   // namespace

auto main(int argc, char** argv) -> int {
    std::println("{}", part_1(argv[1]));
    std::println("{}", part_2(argv[1]));
    return 0;
}
