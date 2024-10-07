import time
import requests

def fetch_data(url):
    print(f"Fetching {url}...")
    response = requests.get(url)  # This blocks the program until the request is complete
    return response.text

def main():
    urls = [
        "https://jsonplaceholder.typicode.com/posts/1",
        "https://jsonplaceholder.typicode.com/posts/2",
        "https://jsonplaceholder.typicode.com/posts/3"
    ]
    
    for url in urls:
        data = fetch_data(url)  # Each call blocks the next one
        print(f"Data fetched from {url}: {data[:60]}...")  # Print first 60 chars

start_time = time.time()
main()
end_time = time.time()
print(f"Total time taken: {end_time - start_time:.2f} seconds")
