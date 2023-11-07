import socket

from _thread import *
import threading

print_lock = threading.Lock()

users:dict = {}

def threadedServ(c: str):
    while True:
        data = c.recv(1024)
        print(data)
        if not data:
            print('[!] Disconected...')

            print_lock.release()
            break
        
        data = data[::-1]
        c.send(data)
    
    c.close()

def Main():
    host = "127.0.0.1"
    port = 9999

    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    print("[+]Conectado a porta", port)

    sock.listen()
    print("[+]Socket ativa")

    while True:
        conn, addr = sock.accept()
        print_lock.acquire()
        print(f"[+]{addr[0]:addr[1]} | conectado")

        start_new_thread(threadedServ, (conn,))
    sock.close()


Main()