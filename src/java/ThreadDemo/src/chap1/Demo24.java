package chap1;

public class Demo24 {
    public static void main(String[] args) {
        System.out.println("主线程的运行优先级是：" + Thread.currentThread().getPriority());
        Thread.currentThread().setPriority(1);
        System.out.println("主线程的运行优先级是：" + Thread.currentThread().getPriority());
        Thread t = new Demo24Thread();
        t.start();
    }
}

class Demo24Thread extends Thread{
    @Override
    public void run() {
        System.out.println("线程的运行优先级是：" + this.getPriority());
    }
}