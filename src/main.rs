use csv;
use reqwest;
use scraper;

struct Course {
    title: Option<String>,
    description: Option<String>,
    prerequisites: Option<Vec<String>>,
    corequisites: Option<Vec<String>>,
}

fn main() {
    let response = reqwest::blocking::get(
        "https://catalog.fullerton.edu/preview_program.php?catoid=80&poid=38156&returnto=11049",
    );
    let html_content = response.unwrap().text().unwrap();

    let document = scraper::Html::parse_document(&html_content);

    println!("{html_content}");
    let html_product_selector = scraper::Selector::parse("li.acalog-course").unwrap();
    let html_courses = document.select(&html_product_selector);

    let mut courses: Vec<Course> = Vec::new();
}
