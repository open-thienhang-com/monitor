import requests
from bs4 import BeautifulSoup
import pandas as pd
import pyarrow.parquet as pq
import pyarrow as pa
import logging
from pmodules.glog import logger
from hdfs import InsecureClient

class WebCrawler:
    def __init__(self, website_url ):
        self.website_url = website_url

    def crawl_data(self):
        try:
            logger.info('crawl_data ...')
            # Send an HTTP GET request to the website
            response = requests.get(self.website_url)
            response.raise_for_status()  # Raise an exception if the request fails

            # Parse the HTML content with BeautifulSoup
            soup = BeautifulSoup(response.text, 'html.parser')
            # Extract and process data from the website (example)
            data = []
            for item in soup.find_all('div', class_=''):
                data.append(item.text)

            return data

        except Exception as e:
            # Log any errors to a log file
            logger.error(f"Error while crawling data: {str(e)}")
            return []

    def save_to_parquet(self, data, parquet_file_path):
        try:
            # Create a Pandas DataFrame from the crawled data
            df = pd.DataFrame({'Data': data})
            arrow_table = pa.Table.from_pandas(df)
            
            # Save the DataFrame to a Parquet file
            pq.write_table(arrow_table, parquet_file_path)

            import os

            current_directory = os.getcwd()
            print("Current Working Directory:", current_directory)

            script_directory = os.path.dirname(os.path.abspath(__file__))
            print("Script Directory (Full Path):", script_directory)

            print(parquet_file_path)

        except Exception as e:
            # Log any errors to a log file
             logger.error(f"Error while saving data to Parquet: {str(e)}")

    def upload_to_hadoop(self):
        try:
            client = InsecureClient('http://localhost:43425/', user='root')
            print("test")
            # Upload the Parquet file to HDFS
            # with open('data.parquet', 'rb') as local_file:
            #     client.upload('a', local_file)
            client.upload('a', 'data.parquet')
        except Exception as e:
            # Log any errors to a log file
            logging.error(f"Error while uploading to Hadoop: {str(e)}")


    def run(self):
        # Create a log file and configure logging
        logger.info('CRAWLER ...')

        # Crawl data from the website
        data = self.crawl_data()

        if data:
            logger.info(data)
            # Save data to a Parquet file
            self.save_to_parquet(data, 'data.parquet')

            # Upload Parquet file to Google Drive
            self.upload_to_hadoop()

        # Print a log message
        logging.info("Data crawling and processing completed.")



