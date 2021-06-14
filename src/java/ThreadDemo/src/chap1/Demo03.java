package chap1;

public class Demo03 {
    public static void main(String[] args) {
        Thread t = new Demo03Thread();
        t.run();
        try {
            for (int i = 0; i < 10; i++) {
                System.out.println("运行main方法");
                Thread.sleep(100);
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}

class Demo03Thread extends Thread{
    @Override
    public void run() {
        try {
            for (int i = 0; i < 10; i++) {
                System.out.println("运行run方法");
                Thread.sleep(100);
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}