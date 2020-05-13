class NodeConfig:

    def __init__(self, node_name: str, isocket_path: str, prometheus_url : str) -> None:
        self.node_name = node_name
        self.isocket_path = isocket_path
        self.prometheus_url = prometheus_url