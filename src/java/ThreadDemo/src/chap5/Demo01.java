package chap5;

public class Demo01 {
    public static void main(String[] args) throws InterruptedException {
        Thread t = new Demo01Thread();
        System.out.println("线程在main方法中的状态1：" + t.getState());
        Thread.sleep(1000);
        t.start();
        Thread.sleep(1000);
        System.out.println("线程在main方法中的状态2：" + t.getState());
    }
}

class Demo01Thread extends Thread{
    public Demo01Thread(){
        System.out.println("构造方法的状态：" + Thread.currentThread().getState());
    }

    @Override
    public void run() {
        System.out.println("run方法的状态：" + Thread.currentThread().getState());
    }
}