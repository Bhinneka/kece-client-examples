import socket

def main():
    client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

    host = '0.0.0.0'
    port = 8000
    client.connect((host, port))

    client.send('SET 1 wury\r\n'.encode())

    response_set = client.recv(4096)

    print(response_set)

    client.send('GET 1\r\n'.encode())

    response_get = client.recv(4096)

    print(response_get)

if __name__ == '__main__':
    main()
