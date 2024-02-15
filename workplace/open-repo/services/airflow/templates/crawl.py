from airflow.decorators import dag, task
from airflow.operators.bash import BashOperator
from pendulum import datetime
from airflow import Dataset
import requests
import uuid
import os

BASE_URL='https://test-adnetwork-core.x.io/openrtb2/auction' #os.getenv(
DATASET_URL = "dataset/device_set_no_1.txt"
def call_auction(dataset):
    id = str(uuid.uuid4())
    payload = {
            "test": 1,
            "id": id,
            "user": {
                "id": "7EAF66B1-04F7-4B90-BA71-B8CE8B59F3D8",
                "buyeruid": ""
            },
            "source": {},
            "device": {
                "ua": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
                "language": "vi",
                "ip": "10.0.10.39",
                "connectiontype": 0,
                "devicetype": 2
            },
            "ext": {},
            "site": {
                "page": "https://demo.x.io/",
                "ref": "https://test-adnetwork-portal.x.io/",
                "search": "?env=test&containerId=cbc89fa4-bcc9-4364-bdc4-2a72c76cea11&inventoryId=837_1593"
            },
            "regs": {
                "ext": {
                    "gdpr": 0
                }
            },
            "imp": [
                {
                    "id": "2a36971a-61f3-4e34-8309-9290d3590711",
                    "native": {
                        "ver": "1.2",
                        "request": "{\"context\":1,\"contextsubtype\":10,\"plcmttype\":1,\"plcmtcnt\":1,\"eventtrackers\":[{\"event\":1,\"methods\":[1,2]}],\"assets\":[{\"id\":0,\"required\":0,\"title\":{\"len\":100000}},{\"id\":1,\"required\":0,\"img\":{\"wmin\":50,\"hmin\":50,\"mimes\":[\"image/jpg\",\"image/gif\",\"image/png\",\"image/webp\"],\"type\":1}},{\"id\":2,\"required\":0,\"data\":{\"type\":2,\"len\":1400000}},{\"id\":3,\"required\":0,\"data\":{\"type\":12,\"len\":2500000}},{\"id\":4,\"required\":0,\"data\":{\"type\":6,\"len\":1000}},{\"id\":5,\"required\":0,\"data\":{\"type\":8,\"len\":9000000}}]}"
                    },
                    "ext": {
                        "prebid": {},
                        "aicadn": {
                            "header_bids": [],
                            "dsps": [],
                            "container_id": 497,
                            "inventory_id": 837,
                            "debug": {
                                "enabled": True,
                                "strategy_id": 1593
                            }
                        },
                        "context": {
                            "path": "/",
                            "referrer": "https://test-adnetwork-portal.x.io/",
                            "search": "?env=test&containerId=cbc89fa4-bcc9-4364-bdc4-2a72c76cea11&inventoryId=837_1593",
                            "title": "x ADS TEST",
                            "url": "https://demo.x.io/?env=test&containerId=cbc89fa4-bcc9-4364-bdc4-2a72c76cea11&inventoryId=837_1593",
                            "sessionId": "",
                            "keywords": "",
                            "description": ""
                        }
                    }
                }
            ]
        }
    response = requests.post(f"{BASE_URL}", json=payload)
    print(response)    

@dag(
    dag_id=dag_id_to_replace,
    start_date=datetime(2023, 7, 1),
    schedule='schedule_to_replace',
    catchup=False,
)
def dag_from_config():
    BashOperator(
        task_id="increase_metrics",
        bash_command='bash_command_to_replace',
        env={"ENVVAR": 'env_var_to_replace'},
        # outlets=[Dataset("dataset/device_set_no_1.txt")],
    )
    @task 
    def load_dataset() -> list:
        path = "dataset/device_set_no_dataset_path.txt"
        # TODO: Check exist device set
        # return exist
        # If no exist call paging api to load device and store back to local cache
        r = requests.get('https://')
        if r.status_code == 200: #todo
            from pathlib import Path
            my_file = Path(path)
            if my_file.is_file() == False:
                try:
                    f = open(path, "a")
                    f.write()
                    f.close()
                except:
                    print('Can not create dataset')
                    #TODO: 
        vdevices = []    
        with open(path, "r") as f:
            contents = f.readlines()
            vdevices.append(contents)

        return [item for sublist in vdevices for item in sublist]
        
    @task
    def run(dataset):
        call_auction(dataset)
        return
    
    # Execute scenarios base on provided dataset
    dataset = load_dataset()
    run(dataset)

dag_from_config()
