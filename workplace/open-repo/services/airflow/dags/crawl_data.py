import datetime
from airflow.decorators import dag, task
from bs4 import BeautifulSoup
import urllib.request as urlrq
import certifi
import ssl
import csv
import subprocess
from hdfs import InsecureClient


@dag(start_date=datetime.datetime.now(), schedule="@daily")
def crawl_data_phone():
    @task(task_id="crawl_data_phone")
    def crawl_data():
        # data_name_product = []
        # data_price_product = []
        # data_list = []
        # for i in range(27):
        #     url = f"https://tiki.vn/dien-thoai-may-tinh-bang/c1789?page={i}"
        #     page = urlrq.urlopen(
        #         url, context=ssl.create_default_context(cafile=certifi.where())
        #     )
        #     soup = BeautifulSoup(page, "html.parser")
        #     name_data = soup.find(
        #         "div", class_="ProductList__NewWrapper-sc-1dl80l2-0 jXFjHV"
        #     ).find_all("div", class_="name")
        #     for name_pro in name_data:
        #         name = name_pro.find("h3")
        #         data_name_product.append(name.text)
        #     price_data = soup.find(
        #         "div", class_="ProductList__NewWrapper-sc-1dl80l2-0 jXFjHV"
        #     ).find_all("div", class_="price-discount__price")
        #     for price in price_data:
        #         data_price_product.append(price.text)
        # for i in range(len(data_name_product) - 1):
        #     temp = {
        #         data_name_product[i],
        #         data_price_product[i],
        #     }
        #     data_list.append(temp)
        # print(data_list)
        client = InsecureClient('http://localhost:9870')
        local_file_path = 'dags/Exported_data.csv'
        hdfs_target_directory = '/Hadoop'
        client.upload(hdfs_target_directory, local_file_path)

        # with open('dags/Exported_data.csv', 'a') as f:
        #     writer = csv.writer(f)
        #     for row in data_list:
        #         writer.writerow(row)
        # command = "hdfs dfs -copyFromLocal dags/Exported_data.csv /Hadoop"
        # result = subprocess.run(command, shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True)

        # # In kết quả
        # print("Standard Output:")
        # print(result.stdout)

        # print("Standard Error:")
        # print(result.stderr)
    crawl_data()


crawl_data_phone()


@dag(start_date=datetime.datetime.now(), schedule="@daily")
def crawl_data_beauti():
    @task(task_id="crawl_data_beauti")
    def crawl_data():
        data_name_product = []
        data_price_product = []
        data_list = []
        for i in range(50):
            url = f"https://tiki.vn/lam-dep-suc-khoe/c1520?page={i}"
            page = urlrq.urlopen(
                url, context=ssl.create_default_context(cafile=certifi.where())
            )
            soup = BeautifulSoup(page, "html.parser")
            name_data = soup.find(
                "div", class_="ProductList__NewWrapper-sc-1dl80l2-0 jXFjHV"
            ).find_all("div", class_="name")
            for name_pro in name_data:
                name = name_pro.find("h3")
                data_name_product.append(name.text)
            price_data = soup.find(
                "div", class_="ProductList__NewWrapper-sc-1dl80l2-0 jXFjHV"
            ).find_all("div", class_="price-discount__price")
            for price in price_data:
                data_price_product.append(price.text)
        for i in range(len(data_name_product) - 1):
            temp = {
                data_name_product[i],
                data_price_product[i],
            }
            data_list.append(temp)

    crawl_data()


crawl_data_beauti()


@dag(start_date=datetime.datetime.now(), schedule="@daily")
def crawl_data_clock_jewelry():
    @task(task_id="crawl_data_clock_jewelry")
    def crawl_data():
        data_name_product = []
        data_price_product = []
        data_list = []
        for i in range(50):
            url = "https://tiki.vn/dong-ho-va-trang-suc/c8371?page={i}"
            page = urlrq.urlopen(
                url, context=ssl.create_default_context(cafile=certifi.where())
            )
            soup = BeautifulSoup(page, "html.parser")
            name_data = soup.find(
                "div", class_="ProductList__NewWrapper-sc-1dl80l2-0 jXFjHV"
            ).find_all("div", class_="name")
            for name_pro in name_data:
                name = name_pro.find("h3")
                data_name_product.append(name.text)
            price_data = soup.find(
                "div", class_="ProductList__NewWrapper-sc-1dl80l2-0 jXFjHV"
            ).find_all("div", class_="price-discount__price")
            for price in price_data:
                data_price_product.append(price.text)
        for i in range(len(data_name_product) - 1):
            temp = {
                data_name_product[i],
                data_price_product[i],
            }
            data_list.append(temp)

    crawl_data()


crawl_data_clock_jewelry()


@dag(start_date=datetime.datetime.now(), schedule="@daily")
def crawl_data_sport():
    @task(task_id="crawl_data_sport")
    def crawl_data():
        data_name_product = []
        data_price_product = []
        data_list = []
        for i in range(50):
            url = f"https://tiki.vn/the-thao-da-ngoai/c1975?page={i}"
            page = urlrq.urlopen(
                url, context=ssl.create_default_context(cafile=certifi.where())
            )
            soup = BeautifulSoup(page, "html.parser")
            name_data = soup.find(
                "div", class_="ProductList__NewWrapper-sc-1dl80l2-0 jXFjHV"
            ).find_all("div", class_="name")
            for name_pro in name_data:
                name = name_pro.find("h3")
                data_name_product.append(name.text)
            price_data = soup.find(
                "div", class_="ProductList__NewWrapper-sc-1dl80l2-0 jXFjHV"
            ).find_all("div", class_="price-discount__price")
            for price in price_data:
                data_price_product.append(price.text)
        for i in range(len(data_name_product) - 1):
            temp = {
                data_name_product[i],
                data_price_product[i],
            }
            data_list.append(temp)

    crawl_data()


crawl_data_sport()


@dag(start_date=datetime.datetime.now(), schedule="@daily")
def crawl_data_laptop():
    @task(task_id="crawl_data_laptop")
    def crawl_data():
        data_name_product = []
        data_price_product = []
        data_list = []
        for i in range(50):
            url = f"https://tiki.vn/laptop-may-vi-tinh-linh-kien/c1846?page={i}"
            page = urlrq.urlopen(
                url, context=ssl.create_default_context(cafile=certifi.where())
            )
            soup = BeautifulSoup(page, "html.parser")
            name_data = soup.find(
                "div", class_="ProductList__NewWrapper-sc-1dl80l2-0 jXFjHV"
            ).find_all("div", class_="name")
            for name_pro in name_data:
                name = name_pro.find("h3")
                data_name_product.append(name.text)
            price_data = soup.find(
                "div", class_="ProductList__NewWrapper-sc-1dl80l2-0 jXFjHV"
            ).find_all("div", class_="price-discount__price")
            for price in price_data:
                data_price_product.append(price.text)
        for i in range(len(data_name_product) - 1):
            temp = {
                data_name_product[i],
                data_price_product[i],
            }
            data_list.append(temp)

    crawl_data()


crawl_data_laptop()


@dag(start_date=datetime.datetime.now(), schedule="@daily")
def crawl_data_house_life():
    @task(task_id="crawl_data_house_life")
    def crawl_data():
        data_name_product = []
        data_price_product = []
        data_list = []
        for i in range(50):
            url = "https://tiki.vn/nha-cua-doi-song/c1883?page={i}"
            page = urlrq.urlopen(
                url, context=ssl.create_default_context(cafile=certifi.where())
            )
            soup = BeautifulSoup(page, "html.parser")
            name_data = soup.find(
                "div", class_="ProductList__NewWrapper-sc-1dl80l2-0 jXFjHV"
            ).find_all("div", class_="name")
            for name_pro in name_data:
                name = name_pro.find("h3")
                data_name_product.append(name.text)
            price_data = soup.find(
                "div", class_="ProductList__NewWrapper-sc-1dl80l2-0 jXFjHV"
            ).find_all("div", class_="price-discount__price")
            for price in price_data:
                data_price_product.append(price.text)
        for i in range(len(data_name_product) - 1):
            temp = {
                data_name_product[i],
                data_price_product[i],
            }
            data_list.append(temp)

    crawl_data()


crawl_data_house_life()


@dag(start_date=datetime.datetime.now(), schedule="@daily")
def crawl_data_balo_vali():
    @task(task_id="crawl_data_balo_vali")
    def crawl_data():
        data_name_product = []
        data_price_product = []
        data_list = []
        for i in range(50):
            url = "https://tiki.vn/balo-va-vali/c6000?page={i}"
            page = urlrq.urlopen(
                url, context=ssl.create_default_context(cafile=certifi.where())
            )
            soup = BeautifulSoup(page, "html.parser")
            name_data = soup.find(
                "div", class_="ProductList__NewWrapper-sc-1dl80l2-0 jXFjHV"
            ).find_all("div", class_="name")
            for name_pro in name_data:
                name = name_pro.find("h3")
                data_name_product.append(name.text)
            price_data = soup.find(
                "div", class_="ProductList__NewWrapper-sc-1dl80l2-0 jXFjHV"
            ).find_all("div", class_="price-discount__price")
            for price in price_data:
                data_price_product.append(price.text)
        for i in range(len(data_name_product) - 1):
            temp = {
                data_name_product[i],
                data_price_product[i],
            }
            data_list.append(temp)

    crawl_data()


crawl_data_balo_vali()


@dag(start_date=datetime.datetime.now(), schedule="@daily")
def crawl_data_store_retail():
    @task(task_id="crawl_data_store_retail")
    def crawl_data():
        data_name_product = []
        data_price_product = []
        data_list = []
        for i in range(50):
            url = "https://tiki.vn/bach-hoa-online/c4384?page={i}"
            page = urlrq.urlopen(
                url, context=ssl.create_default_context(cafile=certifi.where())
            )
            soup = BeautifulSoup(page, "html.parser")
            name_data = soup.find(
                "div", class_="ProductList__NewWrapper-sc-1dl80l2-0 jXFjHV"
            ).find_all("div", class_="name")
            for name_pro in name_data:
                name = name_pro.find("h3")
                data_name_product.append(name.text)
            price_data = soup.find(
                "div", class_="ProductList__NewWrapper-sc-1dl80l2-0 jXFjHV"
            ).find_all("div", class_="price-discount__price")
            for price in price_data:
                data_price_product.append(price.text)
        for i in range(len(data_name_product) - 1):
            temp = {
                data_name_product[i],
                data_price_product[i],
            }
            data_list.append(temp)

    crawl_data()


crawl_data_store_retail()


@dag(start_date=datetime.datetime.now(), schedule="@daily")
def crawl_data_book():
    @task(task_id="crawl_data_book")
    def crawl_data():
        data_name_product = []
        data_price_product = []
        data_list = []
        for i in range(50):
            url = f"https://tiki.vn/nha-sach-tiki/c8322?page={i}"
            page = urlrq.urlopen(
                url, context=ssl.create_default_context(cafile=certifi.where())
            )
            soup = BeautifulSoup(page, "html.parser")
            name_data = soup.find(
                "div", class_="ProductList__NewWrapper-sc-1dl80l2-0 jXFjHV"
            ).find_all("div", class_="name")
            for name_pro in name_data:
                name = name_pro.find("h3")
                data_name_product.append(name.text)
            price_data = soup.find(
                "div", class_="ProductList__NewWrapper-sc-1dl80l2-0 jXFjHV"
            ).find_all("div", class_="price-discount__price")
            for price in price_data:
                data_price_product.append(price.text)

        for i in range(len(data_name_product) - 1):
            temp = {
                data_name_product[i],
                data_price_product[i],
            }
            data_list.append(temp)

    crawl_data()


crawl_data_book()


@dag(start_date=datetime.datetime.now(), schedule="@daily")
def crawl_data_motobike():
    @task(task_id="crawl_data_motobike")
    def crawl_data():
        data_name_product = []
        data_price_product = []
        data_list = []
        for i in range(50):
            url = f"https://tiki.vn/o-to-xe-may-xe-dap/c8594?page={i}"
            page = urlrq.urlopen(
                url, context=ssl.create_default_context(cafile=certifi.where())
            )
            soup = BeautifulSoup(page, "html.parser")
            name_data = soup.find(
                "div", class_="ProductList__NewWrapper-sc-1dl80l2-0 jXFjHV"
            ).find_all("div", class_="name")
            for name_pro in name_data:
                name = name_pro.find("h3")
                data_name_product.append(name.text)
            price_data = soup.find(
                "div", class_="ProductList__NewWrapper-sc-1dl80l2-0 jXFjHV"
            ).find_all("div", class_="price-discount__price")
            for price in price_data:
                data_price_product.append(price.text)
        for i in range(len(data_name_product) - 1):
            temp = {
                data_name_product[i],
                data_price_product[i],
            }
            data_list.append(temp)

    crawl_data()


crawl_data_motobike()


@dag(start_date=datetime.datetime.now(), schedule="@daily")
def crawl_data_motobike():
    @task(task_id="crawl_data")
    def crawl_data():
        data_name_product = []
        data_price_product = []
        data_list = []
        for i in range(50):
            url = f"https://tiki.vn/o-to-xe-may-xe-dap/c8594?page={i}"
            print(url)
            page = urlrq.urlopen(
                url, context=ssl.create_default_context(cafile=certifi.where())
            )
            soup = BeautifulSoup(page, "html.parser")
            name_data = soup.find(
                "div", class_="ProductList__NewWrapper-sc-1dl80l2-0 jXFjHV"
            ).find_all("div", class_="name")
            for name_pro in name_data:
                name = name_pro.find("h3")
                data_name_product.append(name.text)
            print(data_name_product)
            price_data = soup.find(
                "div", class_="ProductList__NewWrapper-sc-1dl80l2-0 jXFjHV"
            ).find_all("span", class_="price-discount__price")
            for price in price_data:
                data_price_product.append(price.text)
        for i in range(len(data_name_product) - 1):
            temp = {data_name_product[i], data_price_product[i]}
            data_list.append(temp)
        print(data_list)

    crawl_data()


crawl_data_motobike()


# Crawl data Satra Food
@dag(start_date=datetime.datetime.now(), schedule="@daily")
def crawl_data_vegetable():
    @task(task_id="crawl_data_vegetable")
    def crawl_data():
        data_name_product = []
        data_price_product = []
        data_list = []
        for i in range(1, 3, 1):
            url = f"https://satrafoods.com.vn/vn/rau-cu/{i}.html"
            print(url)
            page = urlrq.urlopen(
                url, context=ssl.create_default_context(cafile=certifi.where())
            )
            soup = BeautifulSoup(page, "html.parser")
            name_data = soup.find("div", class_="round-item").find_all(
                "div", class_="info"
            )
            for name_pro in name_data:
                name = name_pro.find("h3")
                data_name_product.append(name.text)
            print(data_name_product)
            price_data = soup.find("div", class_="round-item").find_all(
                "span", class_="price"
            )
            for price in price_data:
                data_price_product.append(price.text)
        for i in range(len(data_name_product) - 1):
            temp = {data_name_product[i], data_price_product[i]}
            data_list.append(temp)
        print(data_list)

    crawl_data()


crawl_data_vegetable()


@dag(start_date=datetime.datetime.now(), schedule="@daily")
def crawl_data_meat():
    @task(task_id="crawl_data_meat")
    def crawl_data():
        data_name_product = []
        data_price_product = []
        data_list = []
        for i in range(1, 3, 1):
            url = f"https://satrafoods.com.vn/vn/thit-trung/{i}.html"
            print(url)
            page = urlrq.urlopen(
                url, context=ssl.create_default_context(cafile=certifi.where())
            )
            soup = BeautifulSoup(page, "html.parser")
            name_data = soup.find("div", class_="round-item").find_all(
                "div", class_="info"
            )
            for name_pro in name_data:
                name = name_pro.find("h3")
                data_name_product.append(name.text)
            print(data_name_product)
            price_data = soup.find("div", class_="round-item").find_all(
                "span", class_="price"
            )
            for price in price_data:
                data_price_product.append(price.text)
        for i in range(len(data_name_product) - 1):
            temp = {data_name_product[i], data_price_product[i]}
            data_list.append(temp)
        print(data_list)

    crawl_data()


crawl_data_meat()


@dag(start_date=datetime.datetime.now(), schedule="@daily")
def crawl_data_cool():
    @task(task_id="crawl_data_cool")
    def crawl_data():
        data_name_product = []
        data_price_product = []
        data_list = []
        for i in range(1, 3, 1):
            url = f"https://satrafoods.com.vn/vn/dong-lanh/{i}.html"
            print(url)
            page = urlrq.urlopen(
                url, context=ssl.create_default_context(cafile=certifi.where())
            )
            soup = BeautifulSoup(page, "html.parser")
            name_data = soup.find("div", class_="round-item").find_all(
                "div", class_="info"
            )
            for name_pro in name_data:
                name = name_pro.find("h3")
                data_name_product.append(name.text)
            print(data_name_product)
            price_data = soup.find("div", class_="round-item").find_all(
                "span", class_="price"
            )
            for price in price_data:
                data_price_product.append(price.text)
        for i in range(len(data_name_product) - 1):
            temp = {data_name_product[i], data_price_product[i]}
            data_list.append(temp)
        print(data_list)

    crawl_data()


crawl_data_cool()


@dag(start_date=datetime.datetime.now(), schedule="@daily")
def crawl_data_cool_food():
    @task(task_id="crawl_data_cool_food")
    def crawl_data():
        data_name_product = []
        data_price_product = []
        data_list = []
        for i in range(1, 3, 1):
            url = f"https://satrafoods.com.vn/vn/dong-lanh/{i}.html"
            print(url)
            page = urlrq.urlopen(
                url, context=ssl.create_default_context(cafile=certifi.where())
            )
            print(page)
            soup = BeautifulSoup(page, "html.parser")
            name_data = soup.find("div", class_="round-item").find_all(
                "div", class_="info"
            )
            for name_pro in name_data:
                name = name_pro.find("h3")
                data_name_product.append(name.text)
            print(data_name_product)
            price_data = soup.find("div", class_="round-item").find_all(
                "span", class_="price"
            )
            for price in price_data:
                data_price_product.append(price.text)
        for i in range(len(data_name_product) - 1):
            temp = {data_name_product[i], data_price_product[i]}
            data_list.append(temp)
        print(data_list)

    crawl_data()


crawl_data_cool_food()


@dag(start_date=datetime.datetime.now(), schedule="@daily")
def crawl_data_sweet_food():
    @task(task_id="crawl_data_sweet_food")
    def crawl_data():
        data_name_product = []
        data_price_product = []
        data_list = []
        for i in range(1, 3, 1):
            url = f"https://satrafoods.com.vn/vn/thuc-pham-ngot/{i}.html"
            print(url)
            page = urlrq.urlopen(
                url, context=ssl.create_default_context(cafile=certifi.where())
            )
            soup = BeautifulSoup(page, "html.parser")
            name_data = soup.find("div", class_="round-item").find_all(
                "div", class_="info"
            )
            for name_pro in name_data:
                name = name_pro.find("h3")
                data_name_product.append(name.text)
            print(data_name_product)
            price_data = soup.find("div", class_="round-item").find_all(
                "span", class_="price"
            )
            for price in price_data:
                data_price_product.append(price.text)
        for i in range(len(data_name_product) - 1):
            temp = {data_name_product[i], data_price_product[i]}
            data_list.append(temp)
        print(data_list)

    crawl_data()


crawl_data_sweet_food()
