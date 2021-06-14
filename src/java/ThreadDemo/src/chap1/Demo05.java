package chap1;

public class Demo05 {
    public static void main(String[] args) {
        Runnable r = new Demo05Thread();
        Thread t = new Thread(r);
        System.out.println("运行了main方法");
    }
}

class Demo05Thread implements Runnable{
    @Override
    public void run() {
        System.out.println("运行了run方法");
    }
}