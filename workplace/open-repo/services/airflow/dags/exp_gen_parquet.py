import datetime
from airflow.decorators import dag, task
from airflow.utils.dates import days_ago
from airflow.operators.python import PythonOperator
import pandas as pd
import pyarrow as pa
import pyarrow.parquet as pq

import os
from google.oauth2 import service_account
from googleapiclient.discovery import build
from googleapiclient.http import MediaFileUpload

# Define the default arguments for the DAG
default_args = {
    'owner': 'airflow',
    'start_date': days_ago(1),
    'retries': 1,
}

@dag(
    dag_id='exp_gen_parquet',
    default_args=default_args,
    schedule_interval='@daily',  # Example daily schedule
    catchup=False
)
def exp_gen_parquet():

    @task()
    def crawl_and_store_parquet():
        try:
            data = {
                'Column1': ['value1', 'value2', 'value3'],
                'Column2': ['value1', 'value2', 'value3'],
                # Add more columns as needed
            }

            df = pd.DataFrame(data)
            arrow_table = pa.Table.from_pandas(df)

            # Step 4: Parquet File Storage
            parquet_file_path = 'data/data.parquet'
            pq.write_table(arrow_table, parquet_file_path)

            import os

            current_directory = os.getcwd()
            print("Current Working Directory:", current_directory)

            script_directory = os.path.dirname(os.path.abspath(__file__))
            print("Script Directory (Full Path):", script_directory)

            # ======
            # Set the path to the JSON credentials file you downloaded from the Google Cloud Console
            credentials_path = 'credentials.json'

            # Authenticate with Google Drive API using credentials
            credentials = service_account.Credentials.from_service_account_file(
                credentials_path,
                scopes=['https://www.googleapis.com/auth/drive']
            )

            # Create a Google Drive API service
            drive_service = build('drive', 'v3', credentials=credentials)

            # Specify the file to upload
            file_path = 'file.txt'  # Replace with the path to your file
            file_name = os.path.basename(file_path)

            #
            folder_metadata = {
                'name': 'My Folder',
                'parents': ['1fFRCYdkPxznkX7Kw-Iw3ZwPnjwwzQSQx'],
                'mimeType': 'application/vnd.google-apps.folder'
            }

            created_folder = drive_service.files().create(body=folder_metadata, fields='id').execute()
            print(f"Created folder with ID: {created_folder['id']}")
            # Create a media object for the file
            media = MediaFileUpload(file_path)

            # Create a file on Google Drive
            file_metadata = {
                'name': file_name,
                'parents': ['1fFRCYdkPxznkX7Kw-Iw3ZwPnjwwzQSQx']  # Replace with the ID of the folder where you want to upload the file
            }

            uploaded_file = drive_service.files().create(
                body=file_metadata,
                media_body=media,
                fields='id'
            ).execute()

            print(f'File ID: {uploaded_file["id"]}')

        except Exception as e:
            # Handle the exception here, e.g., log it or take appropriate actions
            print(f"An error occurred: {str(e)}")
    crawl_and_store_parquet()

# Instantiate the DAG
exp_gen_parquet()
