use std::collections::HashMap;

trait Character {
    fn display(&self);
}

struct ConcreteCharacter {
    symbol: char,
}

impl Character for ConcreteCharacter {
    fn display(&self) {
        println!("Character: {}", self.symbol);
    }
}

struct CharacterFactory {
    characters: HashMap<char, Box<dyn Character>>,
}

impl CharacterFactory {
    fn new() -> Self {
        CharacterFactory {
            characters: HashMap::new(),
        }
    }

    fn get_character(&mut self, symbol: char) -> &Box<dyn Character> {
        self.characters
            .entry(symbol)
            .or_insert(Box::new(ConcreteCharacter { symbol }))
    }
}

fn main() {
    let mut character_factory = CharacterFactory::new();

    let char_a = character_factory.get_character('A');
    char_a.display();

    let char_b = character_factory.get_character('B');
    char_b.display();

    character_factory.get_character('A').display();
}
