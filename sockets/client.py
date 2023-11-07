import socket

HOST: str = 'localhost'
PORT: str = 9999

sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

userName: str = input("[🤖]Escolha um nome de usuário: \n")
sock.connect((HOST, PORT))
sock.sendall(str.encode(userName))
while True:
    message = input("[👾]Message: ")
    sock.sendall(str.encode(f"[{userName}] message"))

    data = sock.recv(1024)
    print(data.decode())    