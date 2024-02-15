from pmodules.adapters.http import TestHttpService
from .apiv1 import api_v1

class BaseService:
    def __init__(self):
        service = TestHttpService()
        service.add_api(api_v1.get_api())
        # service.add_api(api_v2.get_api())
        service.run()