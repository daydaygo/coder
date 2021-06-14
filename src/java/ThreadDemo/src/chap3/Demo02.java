package chap3;

public class Demo02 {
    public static void main(String[] args) {
        try {
            /*String str = new String();
            // Exception in thread "main" java.lang.IllegalMonitorStateException
            str.wait();*/

            String str = new String();
            System.out.println("同步代码块之前");
            synchronized (str){
                System.out.println("同步代码块的第一行");
                str.wait();
                System.out.println("wait之后的第一行代码");
            }

        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }
}
