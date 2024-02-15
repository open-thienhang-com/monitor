import os
from pmodules.glog import logger
import subprocess

# proc = subprocess.call("ls", cwd=root_path)
_ROOT_PATH = '/Users/tian/Desktop/workplace/ds_services'
_DEPLOYMENT_PATH = _ROOT_PATH + "/python/deployment"
_ERR_EMPTY = "The input is empty."
_ERR_EXCEPTION = "ERR_CHECKER_001: Can not execute command"


def check_git_version():
    print("OK")    

def get_git_modules(path: str = "") -> str:
    print(path)
    try:
        process = subprocess.Popen(
            [
                "git rev-parse --short HEAD"
            ],
            cwd=path,
            shell=True,
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
            universal_newlines=True
        )
        return process.stdout.readline()
    except:
        return "#ERROR"
    finally:
        print()
        # process.kill()

def get_submodules():
    re_submodules = []
    try:
        process = subprocess.Popen(
            [
                "git config --file .gitmodules --get-regexp path | awk '{ print $2 }'"
            ],
                cwd=_ROOT_PATH,
                shell=True,
                stdout=subprocess.PIPE,
                stderr=subprocess.PIPE,
                universal_newlines=True
            )

        while True:
            name = process.stdout.readline()
            if not name:
                break
            # print([1])
            name = name.replace("\n", "")
            # print()
            git_hash = get_git_modules(_ROOT_PATH+"/"+name)
            re_submodules.append(
                {
                    "name": name.split("python/deployment/",1)[1],
                    "hash": git_hash
                }
            )
            
        #
        subs = []
        for name in os.listdir(_DEPLOYMENT_PATH):
            subs.append(
                {
                    "name": name,
                    "hash": get_git_modules(_DEPLOYMENT_PATH+"/"+name)
                }
            )
        return {
            "root": _ROOT_PATH,
            "working_dir": os.getcwd(),
            "projects": subs,
            "git_modules": re_submodules
        }
    finally:
        print()
        # process.kill()
    
    # subprocess.Popen(["git config --file .gitmodules --get-regexp path | awk '{ print $2 }'"], cwd=deployment_path)
    # update & pull submodule
    # git submodule update --init
    # git submodule update --recursive --remote
    # git pull --recurse-submodules

    # remove submodule
    # result = subprocess.run(
    #     [
    #         # "cd", root_path,
    #         "ls", "-al"
    #     ],
    #     capture_output=True,
    #     text=True
    # ).stdout.strip("\n")
    # print("+++++++++")
    
    
    # print('os.path.abspath(__file__) is:', os.path.abspath(__file__))
    # print('os.path.dirname(os.path.abspath(__file__)) is:', os.path.dirname(os.path.abspath(__file__)))



def add_submodule(gitlab_url: str = "", module_name: str = ""):
    if gitlab_url == "" or module_name == "" :
        raise Exception(_ERR_EMPTY)
    try:
        process = subprocess.Popen(
            [
                "git", "submodule", "add", "--force", gitlab_url, "python/deployment/" + module_name,
                # "git submodule update --init"
            ], 
            cwd=_ROOT_PATH,
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE
        )
        output, err = process.communicate()
        if err != None:
            logger.error(err)
            raise Exception(_ERR_EXCEPTION)
        return output 
    finally:
        process.kill()

def remove_submodule(gitlab_url: str = ""):
    if gitlab_url == "" :
        raise Exception(_ERR_EMPTY)
    try:
        process = subprocess.Popen(
            ["git", "submodule", "add", gitlab_url, "python/deployment"], 
            cwd=_ROOT_PATH,
            # shell=True,
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE
        )
        output, err = process.communicate()
        if err != None:
            logger.error(err)
            raise Exception(_ERR_EXCEPTION)
        return output 
    finally:
        process.kill()
