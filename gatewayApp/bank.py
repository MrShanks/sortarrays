import random


accounts_list = []


class Bank():
    def __init__(self):
        pass

    def create_bank_account(self, name, balance = 0):
        self.name = name
        self.id = generate_id(accounts_list)
        self.balance = balance
        accounts_list.append(self.id)

    def get_info(self):
        data = {'account name' : self.name,
                'account id' : self.id,
                'balance' : self.balance}
        print(data)
        return data


class Account():
    def __init__(self):
        pass

    def get_money(self, account, withdraw):
        account.balance -= withdraw
        print(account.balance)
        return account.balance

    def add_money(self, account, deposit):
        account.balance += deposit
        print(account.balance)
        return account.balance

    def balance(self, account):
        print(account.balance)
        return account.balance


def generate_id(list):
    new_id = random.randrange(100,1000)
    while new_id in list:
        print("already exist")
        new_id = random.randrange(100,1000)
    return new_id



acc1 = Bank()
acc2 = Bank()

acc1.create_bank_account('andrea', 50)
acc2.create_bank_account('simone', 100000)

operator = Account()

acc1.get_info()
acc2.get_info()

operator.balance(acc1)
operator.balance(acc2)
operator.add_money(acc1, 200)
operator.get_money(acc1, 180)