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

    let html_product_selector = scraper::Selector::parse("li.acalog-course").unwrap();
    let html_courses = document.select(&html_product_selector);

    let mut courses: Vec<Course> = Vec::new();

    for html_course in html_courses {
        let title = html_course
            .select(&scraper::Selector::parse("span").unwrap())
            .next()
            .map(|h2| h2.text().collect::<String>());

        let course = Course {
            title,
            description: Some(String::from("My description")),
            prerequisites: Some(Vec::new()),
            corequisites: Some(Vec::new()),
        };

        courses.push(course);
    }

    for course in courses {
        println!("Course Title: {:?}", course.title);
        println!("Course Description: {:?}", course.description);
        println!("Course Prerequisites: {:?}", course.prerequisites);
        println!("Course Corerequisites: {:?}", course.corequisites);
    }
}
