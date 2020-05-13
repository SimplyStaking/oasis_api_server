from configparser import ConfigParser

from setup import setup_user_config_main, setup_user_config_nodes, \
    setup_sentry_config_main

def run() -> None:
    # Initialise parsers
    cp_main = ConfigParser()
    cp_main.read('config/user_config_main.ini')

    cp_nodes = ConfigParser()
    cp_nodes.read('config/user_config_nodes.ini')

    cp_sentry = ConfigParser()
    cp_sentry.read('config/user_config_sentry.ini')

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

        setup_sentry_config_main.setup_nodes(cp_sentry)
        with open('config/user_config_sentry.ini', 'w') as f:
            cp_sentry.write(f)
        print('Saved config/user_config_sentry.ini\n')

        print('Setup completed!')
    except KeyboardInterrupt:
        print('Setup process stopped.')
        return


if __name__ == '__main__':
    run()
