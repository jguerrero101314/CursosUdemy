
import java.util.Scanner;

public class HolaMundo {

    public static void main(String[] args) {
        
        Scanner scanner = new Scanner(System.in);

        System.out.println("proporciona el valor del usuario:");
        var usuario = scanner.nextLine();
        System.out.println("usuario = " + usuario);

    }
}
