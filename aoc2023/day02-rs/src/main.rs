use std::fs;

fn main() {
    let contents = fs::read_to_string(r"../input/02-test.txt")
        .expect("Should have been able to read the file");
    let value: u32 = contents.lines().map(|line: &str| process(line)).sum();
    println!("Part 1 test: {}", value);

    let contents = fs::read_to_string(r"../input/02-prod.txt")
        .expect("Should have been able to read the file");
    let value: u32 = contents.lines().map(|line: &str| process(line)).sum();
    println!("Part 1 prod: {}", value);
}

fn process(line: &str) -> u32 {
    if let Some((game, pulls)) = line.split_once(": ") {
        if pulls
            .split("; ")
            .map(|pull| -> bool { pull.split(", ").map(cubes_are_valid).all(|x| -> bool { x }) })
            .all(|x| -> bool { x })
        {
            return game
                .trim()
                .split_once(' ')
                .expect("Failed split game.")
                .1
                .parse::<u32>()
                .expect("Failed to parse game id.");
        };
    }
    return 0;
}

fn cubes_are_valid(cubes: &str) -> bool {
    match cubes.split_once(' ') {
        Some((count, color)) => {
            let count: u32 = count.parse().expect("Failed to parse count");
            // println!("{color}, {count}");
            match color {
                "red" => count <= 12,
                "green" => count <= 13,
                "blue" => count <= 14,
                _ => panic!("Unknown color!"),
            }
        }
        _ => panic!("Failed to match (count, color)."),
    }
}
