from configparser import ConfigParser
from typing import Optional, List

from setup.utils.config_parsers.sentry import SentryConfig
from setup.utils.user_input import yn_prompt


def get_node(nodes_so_far: List[SentryConfig]) -> Optional[SentryConfig]:
    # Get node's name
    node_names_so_far = [n.node_name for n in nodes_so_far]
    while True:
        node_name = input('Unique node name:\n')
        if node_name in node_names_so_far:
            print('Node name must be unique.')
        else:
            break

    # Get sentry url
    ext_url = input('Sentry Node\'s external url '
                '(typically <IP ADDRESS>:9009):\n')

    # Get tls certificate path
    tls_path = input('Sentry Node\'s tls_identity_cert.pem file location '
                '(typically /serverdir/node/tls_identity.pem):\n')
                
    # Return node
    return SentryConfig(node_name, ext_url, tls_path)


def setup_nodes(cp: ConfigParser) -> None:

    print('==== Sentry')
    print('To retrieve data from Sentry, the API needs'
        'to know the sentry endpoints! The list of '
        'endpoints the API will connect to will now be '
        'set up. Node names must be unique!')

    # Check if list already set up
    already_set_up = len(cp.sections()) > 0
    if already_set_up:
        if not yn_prompt(
                'The list of sentry endpoints is already set up. '
                'Do you wish to replace this list with a new one? (Y/n)\n'):
            return

    # Otherwise ask if they want to set it up
    if not already_set_up and \
            not yn_prompt('Do you wish to set up the '
                'list of sentry endpoints? (Y/n)\n'):
        return

    # Clear config and initialise new list
    cp.clear()
    nodes = []

    # Get node details and append them to the list of nodes
    while True:
        node = get_node(nodes)
        if node is not None:
            nodes.append(node)
            print('Successfully added node.')

        if not yn_prompt('Do you want to add another'
            'sentry endpoint? (Y/n)\n'):
            break

    # Add nodes to config
    for i, node in enumerate(nodes):
        section = 'node_' + str(i)
        cp.add_section(section)
        cp[section]['node_name'] = node.node_name
        cp[section]['ext_url'] = node.ext_url
        cp[section]['tls_path'] = node.tls_path
