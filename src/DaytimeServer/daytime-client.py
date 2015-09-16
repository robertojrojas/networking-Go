import socket
conn=socket.create_connection(('localhost', 1200))
input = conn.makefile()
print input.readline()
conn.close()

# You can use telnet localhost 1200
