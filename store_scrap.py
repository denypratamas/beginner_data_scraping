import pandas as pd
from bs4 import BeautifulSoup
import requests

# Function to product data
def get_store_data(url):
    response = requests.get(url)
    soup = BeautifulSoup(response.text, 'html.parser')
    
    products = soup.find_all('li', class_='product')
    if not products:
        return []

    data = []
    for product in products:
        name = product.find('h2', class_='woocommerce-loop-product__title').text
        price = product.find('span', class_='woocommerce-Price-amount amount').text
        link = product.find('a')['href']
        data.append({'name': name, 'price': price, 'link': link})
    return data

# Function to scrap data from all pages
def get_all_pages(url, max_pages=48):
    all_data = []
    for page_num in range(1, max_pages + 1):
        paged_url = f"{url}/page/{page_num}/"
        print(f"Scraping page {page_num}")
        page_data = get_store_data(paged_url)
        print(f"Products found: {len(page_data)}")
        if not page_data:
            print("No more products found, stopping.")
            break
        all_data.extend(page_data)
    return all_data

# Save it into a CSV file
def save_to_csv(data, filename='store_data.csv'):
    df = pd.DataFrame(data)
    df.to_csv(filename, index=False)
    print(f"Data scraped and saved to {filename}")

def main():
    url = 'https://scrapeme.live/shop/'
    all_data = get_all_pages(url)
    save_to_csv(all_data, 'store_data.csv')

if __name__ == "__main__":
    main()