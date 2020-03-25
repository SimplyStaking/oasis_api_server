from configparser import ConfigParser

from setup import setup_user_config_main, setup_user_config_nodes, setup_prometheus_config_main, setup_extractor_config_main


def run() -> None:
    # Initialise parsers
    cp_main = ConfigParser()
    cp_main.read('config/user_config_main.ini')

    cp_nodes = ConfigParser()
    cp_nodes.read('config/user_config_nodes.ini')
    
    cp_prometheus = ConfigParser()
    cp_prometheus.read('config/prometheus_config_main.ini')
    
    cp_extractor = ConfigParser()
    cp_extractor.read('config/extractor_config_main.ini')
    
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

        setup_extractor_config_main.setup_nodes(cp_extractor)
        with open('config/extractor_config_main.ini', 'w') as f:
            cp_extractor.write(f)
        print('Saved config/extractor_config_main.ini\n')
        
        print('Setup completed!')
    except KeyboardInterrupt:
        print('Setup process stopped.')
        return


if __name__ == '__main__':
    run()
