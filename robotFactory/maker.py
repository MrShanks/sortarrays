import argparse
import yaml
import os
import sys
import factory


# Create the parser
parser = argparse.ArgumentParser(prog = 'Factory Robot',
                                 description='Create robots based on passed order')

# Add an exclusive group to accept a name robot or an order from a file
group = parser.add_mutually_exclusive_group(required=True)

# Add all needed arguments
group.add_argument('-n',
                   '--name',
                   type = str,
                   help = 'name of the robot')

group.add_argument('-o',
                   '--order',
                   type = str,
                   help='Order dict type to process')

parser.add_argument('-t',
                    '--type',
                    choices=['sing', 'fly', 'clean'],
                    type = str,
                    help = 'type of the robot')

# Execute the parse_args() to get passed arguments
args = parser.parse_args()

if args.name != None:
    if args.type != None:
        order_dict = {args.name : args.type}
    else:
        order_dict = {args.name : ''}
else:
    if not os.path.isfile(args.order):
        print('The path specified does not exist')
        sys.exit()
    else:
        with open(args.order) as file:
            order_dict = yaml.load(file, Loader=yaml.FullLoader)

# Create robots based on the passed name/order
for rob in order_dict:
    rob = factory.Factory().create_robot(rob, order_dict[rob])
    print(f'\n{"%06d" % rob.id} - Hi, I call {rob.name.upper()}, Im your new {rob.type.upper()} robot. I also know {rob.drive_ability}.')
    rob.drive()
    rob.main_ability()