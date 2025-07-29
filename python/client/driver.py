

class Driver:

    def __init__(self, url, token, user_agent=""):
        self.config = Configuration()
        self.config.host = url

        if token:
            self.config.api_key["Authorization"] = bearer
            self.config.api_key_prefix["Authorization"] = "Bearer"
            self.client = ApiClient(configuration=self.config)

