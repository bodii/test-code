trait UIElement {
    fn render(&self);
}

struct Button;

impl UIElement for Button {
    fn render(&self) {
        println!("Rendered a button");
    }
}

struct Checkbox;

impl UIElement for Checkbox {
    fn render(&self) {
        println!("Rendered a Checkbox");
    }
}

trait UIFactory {
    fn create_button(&self) -> Box<dyn UIElement>;
    fn create_checkbox(&self) -> Box<dyn UIElement>;
}

struct WindowsUIFactory;

impl UIFactory for WindowsUIFactory {
    fn create_button(&self) -> Box<dyn UIElement> {
        Box::new(Button)
    }

    fn create_checkbox(&self) -> Box<dyn UIElement> {
        Box::new(Checkbox)
    }
}

fn main() {
    let factory = WindowsUIFactory;
    let button = factory.create_button();
    let checkout = factory.create_checkbox();

    button.render();
    checkout.render();
}
