package chap2;

public class Demo26 {
    public static void main(String[] args) throws InterruptedException {
        Demo26Service service = new Demo26Service();
//        service.foo();
        service.start();
        Thread.sleep(100);
        System.out.println(Thread.currentThread().getName() + "准备停止foo方法的循环");
        service.flag = false;
    }
}

class Demo26Service extends Thread{
    // 标志，控制循环
    volatile public boolean flag = true;

    public void run(){
        foo();
    }

    public void foo(){
        System.out.println("foo开始运行");
        while(flag) {

        }
        System.out.println("foo运行结束");
    }
}