use std::fs;

#[derive(Clone, Debug)]
struct Mapping(String, char);

fn main() {
    let mapping = vec![
        Mapping(String::from("1"), '1'),
        Mapping(String::from("2"), '2'),
        Mapping(String::from("3"), '3'),
        Mapping(String::from("4"), '4'),
        Mapping(String::from("5"), '5'),
        Mapping(String::from("6"), '6'),
        Mapping(String::from("7"), '7'),
        Mapping(String::from("8"), '8'),
        Mapping(String::from("9"), '9'),
        Mapping(String::from("one"), '1'),
        Mapping(String::from("two"), '2'),
        Mapping(String::from("three"), '3'),
        Mapping(String::from("four"), '4'),
        Mapping(String::from("five"), '5'),
        Mapping(String::from("six"), '6'),
        Mapping(String::from("seven"), '7'),
        Mapping(String::from("eight"), '8'),
        Mapping(String::from("nine"), '9'),
    ];
    let mut reverse_mapping = Vec::<Mapping>::new();
    for Mapping(key, value) in mapping.iter() {
        reverse_mapping.push(Mapping(key.chars().rev().collect(), *value));
    }

    let contents = fs::read_to_string(r"../input/01-input.txt")
        .expect("Should have been able to read the file");
    let value: u32 = contents
        .lines()
        .map(|line: &str| process(&String::from(line), &mapping, &reverse_mapping))
        .sum();
    println!("Result for part 2: {}", value);
}

fn process(line: &String, mapping: &Vec<Mapping>, reverse_mapping: &Vec<Mapping>) -> u32 {
    let mut result = String::new();
    result.push(find_digit(line, mapping).expect("No digit was found."));
    result.push(
        find_digit(&line.chars().rev().collect(), &reverse_mapping)
            .expect("No digit was found (reverse)."),
    );

    return result.parse::<u32>().unwrap();
}

fn find_digit(line: &String, mapping: &Vec<Mapping>) -> Option<char> {
    for i in 0..line.len() {
        for Mapping(key, value) in mapping {
            if line[i..].starts_with(key) {
                return Some(*value);
            }
        }
    }
    None
}
