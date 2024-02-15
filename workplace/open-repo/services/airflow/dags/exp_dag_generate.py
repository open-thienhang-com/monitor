from airflow.operators.bash import BashOperator
from pendulum import datetime
from airflow import Dataset
from airflow.decorators import dag, task
from pathlib import Path
import json, requests
import os
import shutil
import fileinput
from requests import Request, Session
from textwrap import dedent

dirname=os.path.dirname
filepath = dirname(dirname(os.path.abspath(__file__)))

BASE_URL='http://host.docker.internal:8000' #os.getenv(

@dag(
    start_date=datetime(2022, 10, 1),
    schedule=None,
    catchup=False
)
def adx_generate_auction():
    @task()
    def getBasicToken() -> str:
        return "BUG" #TODO
    
    @task()
    def getScenario(access_token: str):
        print("******")
        print(access_token)
        s = Session()
        headers = {
            "Content-Type": "application/json; charset=utf-8",
            "authorization": "Bearer " + access_token
        }
        r = Request('GET', f"{BASE_URL}/v1/scenario", data={}, headers=headers)
        prepped = r.prepare()

        jobs = []
        #TODO: hardcode get maximum 100 active jobs, need paging
        for _ in range(5):
            # prepped.headers['Content-Type'] Add page / offset
            resp = s.send(prepped,  timeout=300)
            if resp.status_code == 200:
                resp = resp.json()
                j = resp['data']
                if len(j) == 0:
                    continue
                jobs = jobs + j
            else:
                continue
        return jobs
    
    @task
    def generate_auction_scenario(jobs: any):
        print(jobs)
        for j in jobs:
            config = {
                "dag_id": j["dag_id"],
                "schedule": j["schedule"],
                "bash_command": j["bash_command"],
                "env_var": j["env_var_to"],
                "dataset_path": str(j["set_id"])
            }

            new_filename = filepath + "/dags/adx-auction_{}.py".format(str(config["dag_id"]))
            shutil.copyfile(filepath+ "/templates/adx-auction.py", new_filename)

            for line in fileinput.input(new_filename, inplace=True ):
                line = line.replace("dag_id_to_replace", "'auction_dataset_no_" + config["dag_id"] + "'")
                line = line.replace("schedule_to_replace", config["schedule"])
                line = line.replace("bash_command_to_replace", config["bash_command"])
                line = line.replace("env_var_to_replace", config["env_var"])
                line = line.replace("dataset_path", config["dataset_path"])
                print(line, end="")
    
    token = getBasicToken()
    jobs = getScenario(token)
    generate_auction_scenario(jobs)

adx_generate_auction()
