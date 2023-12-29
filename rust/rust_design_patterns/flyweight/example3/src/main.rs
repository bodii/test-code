use std::collections::HashMap;

trait Font {
    fn print(&self);
}

struct ConcreteFont {
    name: String,
    size: u32,
}

impl Font for ConcreteFont {
    fn print(&self) {
        println!("Font: {}, Size: {}", self.name, self.size);
    }
}

struct FontFactory {
    fonts: HashMap<String, Box<dyn Font>>,
}

impl FontFactory {
    fn new() -> Self {
        FontFactory {
            fonts: HashMap::new(),
        }
    }

    fn get_font(&mut self, name: &str, size: u32) -> &Box<dyn Font> {
        let key = format!("{}..{}", name, size);
        self.fonts
            .entry(key.clone())
            .or_insert(Box::new(ConcreteFont {
                name: name.to_string(),
                size,
            }))
    }
}

fn main() {
    let mut font_factory = FontFactory::new();

    let arial_12 = font_factory.get_font("Arial", 12);
    arial_12.print();

    let times_14 = font_factory.get_font("Times New Roman", 14);
    times_14.print();

    let arial_12_2 = font_factory.get_font("Arial", 12);
    arial_12_2.print();
}
