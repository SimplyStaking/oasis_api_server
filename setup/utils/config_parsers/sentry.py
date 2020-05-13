class SentryConfig:

    def __init__(self, node_name: str, ext_url: str, tls_path: str) -> None:
        self.node_name = node_name
        self.ext_url = ext_url
        self.tls_path = tls_path

