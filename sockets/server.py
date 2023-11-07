import socket

HOST: str = 'localhost'
PORT: str = 9999

sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

sock.bind((HOST, PORT))

sock.listen()
print(f"[+]Servidor {HOST}:{PORT} aguardando conexão")

users: dict = {}

while True:
    # conn == client socket, addr == client port
    conn, addr = sock.accept()
    print(f"[+]Conexão de {addr}")
    
    #serverside username
    
    # data = conn.recv(1024).decode()
    # users[addr] = data
    while True:
        data = conn.recv(1024).decode()
        print(data)
        conn.sendall(str.encode(f"[{users[addr]}] {data}"))
        
        if data == "/die":
            conn.sendall(str.encode("vsf irmao nem queria vc aqui na minha socket mesmo"))
            conn.close()
            print(f"[💀]{addr} is dead")