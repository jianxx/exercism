import threading

class BankAccount:
    def __init__(self):
        self.opened = False
        self.balance = 0
        self.lock = threading.Lock()

    def get_balance(self):
        if not self.opened:
            raise ValueError("Account is not open")
        return self.balance

    def open(self):
        if self.opened:
            raise ValueError("Do not open an existing account")
        self.opened = True

    def deposit(self, amount):
        if not self.opened or amount < 0:
            raise ValueError("Account is not open")
        else:
            self.__add_to_balance(amount)

    def withdraw(self, amount):
        if not self.opened or amount > self.balance or amount < 0:
            raise ValueError("Account is not open")
        else:
            self.__add_to_balance(0-amount)

    def close(self):
        if not self.opened:
            raise ValueError("Account is not open")
        self.opened = False
        self.balance = 0

    def __add_to_balance(self, amount):
        self.lock.acquire()
        newbalance = self.balance + amount
        if newbalance >= 0:
            self.balance = newbalance
            self.lock.release()
        else:
            self.lock.release()
            raise ValueError("New balance < 0")
