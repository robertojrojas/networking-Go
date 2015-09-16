require 'socket'

hostname = 'localhost'
port = 1200

s = TCPSocket.open(hostname, port)

while line = s.gets  # Read line from the socket  
  puts line.chomp 
end 
s.close 
