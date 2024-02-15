import os
import sys
current_path = os.getcwd()
print("Current Working Directory:", current_path)

script_path = os.path.abspath(__file__)
print("Script Location:", script_path)


from services.http import BaseService

if __name__ == "__main__":
    BaseService()
