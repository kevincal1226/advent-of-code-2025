import sys
from functools import reduce
import math

import heapq
from typing import List

K = 1000


class Edge:
    def __init__(self, a: int, a_pts: List, b: int, b_pts: List, dist: float) -> None:
        self.a = a
        self.a_pts = a_pts
        self.b_pts = b_pts
        self.b = b
        self.dist = dist

    def __lt__(self, other):
        return self.dist > other.dist


def find(a: int, reps: List) -> int:
    if reps[a] != a:
        reps[a] = find(reps[a], reps)

    return reps[a]


def unionize(a: int, b: int, reps: List):
    reps[find(a, reps)] = reps[find(b, reps)]


def get_input(filename: str) -> List:
    points = []
    with open(filename) as f:
        lines = f.readlines()
        for line in lines:
            (
                x,
                y,
                z,
            ) = line.split(",")
            points.append([int(x), int(y), int(z)])

    return points


def part1(filename: str):
    points = get_input(filename)

    smallest_edges = []
    for a_idx, a in enumerate(points):
        for b_idx, b in enumerate(points):
            if b_idx <= a_idx:
                continue

            dist = math.sqrt(
                reduce(lambda acc, x: acc + (x[0] - x[1]) ** 2, zip(a, b), 0)
            )

            if smallest_edges and smallest_edges[0].dist > dist:
                heapq.heappop(smallest_edges)

            if len(smallest_edges) < K:
                heapq.heappush(smallest_edges, Edge(a_idx, a, b_idx, b, dist))

    reps = [i for i in range(len(points))]
    smallest_edges.sort(reverse=True)

    for edge in smallest_edges:
        unionize(min(edge.a, edge.b), max(edge.a, edge.b), reps)

    repcounts = [0 for _ in range(len(points))]

    for rep in reps:
        repcounts[find(rep, reps)] += 1

    repcounts.sort(reverse=True)

    print(repcounts[0] * repcounts[1] * repcounts[2])


def part2(filename: str):
    points = get_input(filename)

    edges: List[Edge] = []
    for a_idx, a in enumerate(points):
        for b_idx, b in enumerate(points):
            if b_idx <= a_idx:
                continue

            dist = math.sqrt(
                reduce(lambda acc, x: acc + (x[0] - x[1]) ** 2, zip(a, b), 0)
            )

            edges.append(Edge(a_idx, a, b_idx, b, dist))

    edges.sort(reverse=True)

    reps = [i for i in range(len(points))]
    for edge in edges:
        unionize(min(edge.a, edge.b), max(edge.a, edge.b), reps)

        if all(find(i, reps) == find(reps[0], reps) for i in reps):
            print(edge.a_pts[0] * edge.b_pts[0])
            return


def main():
    part1(sys.argv[1])
    part2(sys.argv[1])


main()
