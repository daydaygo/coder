package chap1;

public class Demo01 {
    public static void main(String[] args) throws InterruptedException {
        while(true){
            System.out.println(System.currentTimeMillis());
            Thread.sleep(1000);
        }
    }
}
