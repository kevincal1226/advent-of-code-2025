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

#include <boost/algorithm/string.hpp>

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

    return std::ranges::fold_left(input, 0ULL, [](uint64_t acc, Column& c) {
        return acc
             + std::ranges::fold_left(c.nums, c.operation == '*' ? 1ULL : 0ULL, [&](uint64_t sub_acc, uint64_t n) {
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
    std::vector<std::string> verticals;

    std::string line;
    while (std::getline(ifs, line)) {
        if (verticals.empty()) {
            verticals.resize(line.size());
        }

        for (const auto& [c, word] : std::ranges::zip_view(line, verticals)) {
            word.push_back(c);
        }
    }

    std::ranges::for_each(verticals, [&](std::string& vert) {
        boost::algorithm::trim(vert);
        if (vert.empty()) {
            return;
        }

        if (vert.back() == '*' || vert.back() == '+') {
            input.emplace_back();
            input.back().operation = vert.back();
            vert.pop_back();
        }

        input.back().nums.emplace_back(std::stoull(vert));
    });

    return std::ranges::fold_left(input, 0ULL, [](uint64_t acc, Column& c) {
        return acc
             + std::ranges::fold_left(c.nums, c.operation == '*' ? 1ULL : 0ULL, [&](uint64_t sub_acc, uint64_t n) {
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
