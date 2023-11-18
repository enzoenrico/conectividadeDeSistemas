
import socket
import threading
def recebeDados(conn,ender):

try:
nome = conn.recv(50).decode() 
print(f"Conectado com {nome}, IP: {ender[0]}, PORTA: {ender[1]}")
except:
print("Ocorreu um erro durante o recebimento do nome. A conexão será
encerrada")
return
while True:
try:
mensagem = conn.recv(1024).decode()
mensagemNome = nome + " >> " + mensagem
print(mensagemNome)

broadcast)
try:
for cliente in lista_cliente:
cliente.sendall(mensagemNome.encode())
except:
print("Ocorreu algum erro no envio dos dados...")
return
except:
print("Ocorreu algum erro na recepção dos dados, encerrando conexão.")
break
HOST = 'localhost' 
PORT = 9999
lista_cliente = []



sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)


sock.bind((HOST,PORT))

sock.listen() 
print(f"O Servidor {HOST}:{PORT} está aguardando conexões")
while True:



try:
conn, ender = sock.accept()

lista_cliente.append(conn)
except:
print('Ocorreu um erro durante o ACCEPT() na conexão com um novo usuário')
continue

threadCliente = threading.Thread(target=recebeDados, args=[conn,ender])
threadCliente.start()
