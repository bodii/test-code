use crate::gui::Dialog;
use crate::html_gui::HtmlDialog;
use crate::windows_gui::WindowsDialog;

pub fn initialize() -> &'static dyn Dialog {
    if cfg!(windows) {
        println!("-- Windows detected, creating Windows GUI --");
        &HtmlDialog
    } else {
        println!("-- No OS detected, creating the HTML GUI --");
        &WindowsDialog
    }
}