import socket
import threading

def receiveData(conn, addr):
    try:
        #buffer size 50
        name = conn.recv(50).decode() 
    except:
        #connect not ok
        return
    while True:
        try:
            #get message from client
            data = conn.recv(1024).decode()
            msg = f"[{name}]: {data}"
            try:
                for client in clients:
                    client.sendall(msg.encode())
            except:
                #couldn't send message
                return
        except:
            #couldn't recieve data
            break

HOST = 'localhost' 
PORT = 9999
clients = []

sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

sock.bind((HOST,PORT))

sock.listen()
print("[+]Server listening")
while True:
    try:
        conn, addr = sock.accept()

        clients.append(conn)
    except:
        #couldn't accept connection
        continue
    
    #start thread to recieve data
    threadCliente = threading.Thread(target=receiveData, args=[conn,addr])
    threadCliente.start()
    
  
     
   

