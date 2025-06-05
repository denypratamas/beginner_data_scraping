🛒 Web Scraper for WooCommerce Store
This project is a Python web scraper built to extract product data (name, price, and product link) from a WooCommerce-based store — specifically from https://scrapeme.live/shop, a public dummy store used for scraping practice.

🚀 Features
Scrapes product data from multiple pages.
Extracts product name, price, and direct link.
Saves the data into a CSV file.
Built with requests, BeautifulSoup, and pandas.

📦 Output
Scraped data is saved as a CSV file:

Sample Columns:
name — Product title
price — Price in string format (includes currency)
link — Direct link to the product page


🔧 Requirements
Python 3.x

Packages:
pandas
requests
beautifulsoup4


🧠 How It Works
get_store_data(url)
Scrapes a single page of product listings.

get_all_pages(base_url, max_pages=48)
Iterates through paginated URLs until no products are found or max page is reached.

save_to_csv(data, filename)
Saves the collected data to a CSV file.

main()
Entry point for running the scraper.


▶️ Usage
Just run:

bash
Copy
Edit
py scraper.py
This will scrape the data and create a file named store_data.csv in the current directory.

📌 Notes
Target site: https://scrapeme.live/shop
This site is intentionally made for scraping practice — no ethical or legal issues involved.
Scraper stops automatically if no products are found on a page.

✍️ Author
Made with 💻 by denypratamas

📄 License
This project is open source and available under the MIT License.
