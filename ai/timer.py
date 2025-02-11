#!/bin/env python3
"""
This is CANITEY MODIFICATION
GOAL: ADD A TIMER COUNT THAT WILL FIRE EVENT/ACITON IF THE SPAN IS EXCEEDED
"""
import time


class Timer:
    def __init__(self):
        self.now__ = time.time()

    def start(self):
        current = time.time()
        if current - self.now__ >= 3:
            print("Nooo")
        

def main():
    while True:
        timer = Timer()
        timer.start()


if __name__ == "__main__":
    main()
