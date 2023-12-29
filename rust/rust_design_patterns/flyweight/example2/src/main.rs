use std::collections::HashMap;

trait Icon {
    fn display(&self, x: i32, y: i32);
}

struct ConcreteIcon {
    filename: String,
}

impl Icon for ConcreteIcon {
    fn display(&self, x: i32, y: i32) {
        println!("Displaying {} at position ({}, {})", self.filename, x, y);
    }
}

struct IconFactory {
    icons: HashMap<String, Box<dyn Icon>>,
}

impl IconFactory {
    fn new() -> Self {
        IconFactory {
            icons: HashMap::new(),
        }
    }

    fn get_icon(&mut self, filename: &str) -> &Box<dyn Icon> {
        self.icons
            .entry(filename.to_string())
            .or_insert(Box::new(ConcreteIcon {
                filename: filename.to_string(),
            }))
    }
}

fn main() {
    let mut icon_factory = IconFactory::new();

    icon_factory.get_icon("document.png").display(10, 10);

    let folder_icon = icon_factory.get_icon("folder.png");
    folder_icon.display(20, 20);

    let document_icon2 = icon_factory.get_icon("document.png");
    document_icon2.display(30, 30);
}
