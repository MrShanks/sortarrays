import factory

order = {
        'batman' : 'sing',
        'superman' : 'fly',
        'spalman' : 'clean',
        'doraemon' : ''
}


for rob in order:
    rob = factory.Factory().create_robot(rob, order[rob])
    print(f'\n{rob.id} - Hi, I call {rob.name.upper()}, Im your new {rob.type.upper()} robot. I also know {rob.drive_ability}.')
    rob.drive()
    rob.main_ability()