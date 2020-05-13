from configparser import ConfigParser
from typing import Optional, List

from setup.utils.config_parsers.node import NodeConfig
from setup.utils.user_input import yn_prompt


def get_node(nodes_so_far: List[NodeConfig]) -> Optional[NodeConfig]:
    
    # Get node's name
    node_names_so_far = [n.node_name for n in nodes_so_far]
    while True:
        node_name = input('Unique node name:\n')
        if node_name in node_names_so_far:
            print('Node name must be unique.')
        else:
            break

    # Get node's internal socket path way
    isocket_path = input('Node\'s internal socket file path which '
                    'was setup during the node\'s installation '
                    '(typically unix:/serverdir/nodes/internal.sock):\n')

    print('==== Prometheus')
    print('To retrieve data from Prometheus, the API server needs'
          'to have the Prometheus endpoints! ')

    # Get Prometheus's URL of the Node
    prometheus_url = input('Prometheus Node\'s localhost url '
                '(typically http://127.0.0.1:3000):\n')

    # Return node
    return NodeConfig(node_name, isocket_path, prometheus_url)


def setup_nodes(cp: ConfigParser) -> None:

    print('==== Nodes')
    print('To retrieve data from nodes, the API needs to know where to find '
          'the nodes! The list of nodes the API will connect to will now be '
          'set up. This includes validators, sentries, and any full nodes that '
          'can be used as a data source to retrieve data from the network\'s '
          'perspective. Node names must be unique! The list of API nodes must '
          'also include their Prometheus endpoints.')

    # Check if list already set up
    already_set_up = len(cp.sections()) > 0
    if already_set_up:
        if not yn_prompt(
                'The list of nodes is already set up. Do you wish '
                'to replace this list with a new one? (Y/n)\n'):
            return

    # Otherwise ask if they want to set it up
    if not already_set_up and \
            not yn_prompt('Do you wish to set up the list of nodes? (Y/n)\n'):
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

        if not yn_prompt('Do you want to add another node? (Y/n)\n'):
            break

    # Add nodes to config
    for i, node in enumerate(nodes):
        section = 'node_' + str(i)
        cp.add_section(section)
        cp[section]['node_name'] = node.node_name
        cp[section]['isocket_path'] = node.isocket_path
        cp[section]['prometheus_url'] = node.prometheus_url
