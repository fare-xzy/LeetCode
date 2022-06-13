import time
from multiprocessing import Process, Queue, Pool


def processTest(n):
    print(n * n)


def processPool():
    p = Pool(processes=10)  # 进程池

if __name__ == "__main__":
    q = Queue()
    for i in range(100):
        p = Process(target=processTest, args=(i,))
        p.start()
        # p.join() # 强制等待
        # p.terminate()  # 终止
