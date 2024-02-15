import os
import sys
print("helloworld")
print(os.path.join(os.path.dirname(__file__)))

from pmodules.crawler import WebCrawler

class Service:
    def __init__(self):
        # Example usage:
        website_url = 'https://en.wikipedia.org/wiki/Da_Lat'

        crawler = WebCrawler(website_url)
        crawler.run()
        