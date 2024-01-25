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
    // for html_product in html_products {
    //     // scraping logic to retrieve the info
    //     // of interest
    //     let url = html_product
    //         .select(&scraper::Selector::parse("a").unwrap())
    //         .next()
    //         .and_then(|a| a.value().attr("href"))
    //         .map(str::to_owned);
    //     let image = html_product
    //         .select(&scraper::Selector::parse("img").unwrap())
    //         .next()
    //         .and_then(|img| img.value().attr("src"))
    //         .map(str::to_owned);
    //     let name = html_product
    //         .select(&scraper::Selector::parse("h2").unwrap())
    //         .next()
    //         .map(|h2| h2.text().collect::<String>());
    //     let price = html_product
    //         .select(&scraper::Selector::parse(".price").unwrap())
    //         .next()
    //         .map(|price| price.text().collect::<String>());

    //     // instantiate a new Pokemon product
    //     // with the scraped data and add it to the list
    //     let pokemon_product = PokemonProduct {
    //         url,
    //         image,
    //         name,
    //         price,
    //     };
    //     pokemon_products.push(pokemon_product);
    // }

    // // create the CSV output file
    // let path = std::path::Path::new("products.csv");
    // let mut writer = csv::Writer::from_path(path).unwrap();

    // // append the header to the CSV
    // writer
    //     .write_record(&["url", "image", "name", "price"])
    //     .unwrap();
    // // populate the output file
    // for product in pokemon_products {
    //     let url = product.url.unwrap();
    //     let image = product.image.unwrap();
    //     let name = product.name.unwrap();
    //     let price = product.price.unwrap();
    //     writer.write_record(&[url, image, name, price]).unwrap();
    // }

    // // free up the resources
    // writer.flush().unwrap();
}
