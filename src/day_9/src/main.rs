use std::{
    cmp::{max, min},
    fs::File,
    io::{BufRead, BufReader},
};

use geo::{Covers, Polygon};

fn get_input(file: String) -> Vec<(i64, i64)> {
    let reader = BufReader::new(File::open(file).unwrap());
    reader
        .lines()
        .map(|line| line.unwrap())
        .map(|line| {
            let (y, x) = line.split_once(",").unwrap();
            (y.parse::<i64>().unwrap(), x.parse::<i64>().unwrap())
        })
        .collect()
}

fn part1(file: String) -> i64 {
    let points = get_input(file);
    points
        .iter()
        .enumerate()
        .map(|(i, (y1, x1))| {
            points
                .iter()
                .skip(i + 1)
                .map(|(y2, x2)| (max(y1, y2) - min(y1, y2) + 1) * (max(x1, x2) - min(x1, x2) + 1))
                .max()
                .unwrap_or(0)
        })
        .max()
        .unwrap()
}

fn part2(file: String) -> i64 {
    let mut points: Vec<(f64, f64)> = get_input(file)
        .iter()
        .map(|p| (p.0 as f64, p.1 as f64))
        .collect();
    points.push(*points.first().unwrap());

    let outer_poly = Polygon::new(points.clone().into(), vec![]);

    let mut best_area = 0;

    let look_below: (i64, i64) = (94916, 50260);
    let look_above: (i64, i64) = (94916, 48492);

    let mut point_pairs: Vec<((i64, i64), (i64, i64))> = points
        .into_iter()
        .map(|p1| {
            if p1.1 as i64 <= look_above.1 {
                (look_above, (p1.0 as i64, p1.1 as i64))
            } else {
                (look_below, (p1.0 as i64, p1.1 as i64))
            }
        })
        .collect();

    point_pairs.sort_by_key(|((y1, x1), (y2, x2))| {
        -(max(y1, y2) - min(y1, y2) + 1) * (max(x1, x2) - min(x1, x2) + 1)
    });

    println!("Done sorting");

    point_pairs.into_iter().for_each(|((y1, x1), (y2, x2))| {
        let curr_area = (max(y1, y2) - min(y1, y2) + 1) * (max(x1, x2) - min(x1, x2) + 1);

        if curr_area > best_area {
            let x_min = x1.min(x2) as f64;
            let x_max = x1.max(x2) as f64;
            let y_min = y1.min(y2) as f64;
            let y_max = y1.max(y2) as f64;
            let a = vec![
                (y_min, x_min),
                (y_min, x_max),
                (y_max, x_max),
                (y_max, x_min),
                (y_min, x_min),
            ];
            let inner = Polygon::new(a.into(), vec![]);
            if outer_poly.covers(&inner) {
                best_area = curr_area;
            }
        }
    });

    best_area
}

fn main() {
    println!("{}", part1(std::env::args().nth(1).unwrap()));
    println!("{}", part2(std::env::args().nth(1).unwrap()))
}
