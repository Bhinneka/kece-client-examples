
import java.io.BufferedReader;
import java.io.DataOutputStream;
import java.io.IOException;
import java.io.InputStreamReader;
import java.net.Socket;

/**
 *
 * @author wurianto
 */
public class Main {
    
    private Socket client;
    private DataOutputStream outputStream;
    private BufferedReader reader;
    
    public void startConn(String host, int port) throws IOException {
        client = new Socket(host, port);
        outputStream = new DataOutputStream(client.getOutputStream());
        reader = new BufferedReader(new InputStreamReader(client.getInputStream()));
    }
    
    public String sendMessage(byte[] message) throws IOException {
        this.outputStream.write(message);
        String reply = reader.readLine();
        return reply;
        
    }
    
    public void close() throws IOException {
        client.close();
        outputStream.close();
        reader.close();
    }
    
    public static void main(String[] args) throws IOException {
        Main s = new Main();
        s.startConn("localhost", 8000);
        s.sendMessage("SET 1 wury\r\n".getBytes());
        String reply = s.sendMessage("GET 1\r\n".getBytes());
        System.out.println(reply);
        s.close();
    }
    
}
