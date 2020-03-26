from configparser import ConfigParser

from setup import setup_user_config_main, setup_user_config_nodes
from setup import setup_prometheus_config_main, setup_node_exporter_nodes

def run() -> None:
    # Initialise parsers
    cp_main = ConfigParser()
    cp_main.read('config/user_config_main.ini')

    cp_nodes = ConfigParser()
    cp_nodes.read('config/user_config_nodes.ini')
    
    cp_prometheus = ConfigParser()
    cp_prometheus.read('config/prometheus_config_main.ini')
    
    cp_exporter = ConfigParser()
    cp_exporter.read('config/node_exporter_nodes.ini')
    
    # Start setup
    print('Welcome to the Oasis API Server setup script!')
    try:
        setup_user_config_main.setup_all(cp_main)
        with open('config/user_config_main.ini', 'w') as f:
            cp_main.write(f)
        print('Saved config/user_config_main.ini\n')

        setup_user_config_nodes.setup_nodes(cp_nodes)
        with open('config/user_config_nodes.ini', 'w') as f:
            cp_nodes.write(f)
        print('Saved config/user_config_nodes.ini\n')

        setup_prometheus_config_main.setup_nodes(cp_prometheus)
        with open('config/prometheus_config_main.ini', 'w') as f:
            cp_prometheus.write(f)
        print('Saved config/prometheus_config_main.ini\n')

        setup_node_exporter_nodes.setup_nodes(cp_exporter)
        with open('config/node_exporter_nodes.ini', 'w') as f:
            cp_exporter.write(f)
        print('Saved config/node_exporter_nodes.ini\n')
        
        print('Setup completed!')
    except KeyboardInterrupt:
        print('Setup process stopped.')
        return


if __name__ == '__main__':
    run()
