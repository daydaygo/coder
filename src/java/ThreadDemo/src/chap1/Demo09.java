package chap1;

public class Demo09 {
    public static void main(String[] args) {
        Thread t = new Demo09Thread();
        System.out.println("准备开始启动线程：" + t.isAlive());
        t.start();  // 启动线程
        System.out.println("启动线程后：" + t.isAlive());
    }
}

class Demo09Thread extends Thread{
    public Demo09Thread(){
        System.out.println("构造方法开始");
        System.out.println(Thread.currentThread().isAlive());
        System.out.println(this.isAlive());
        System.out.println("构造方法结束");
    }

    @Override
    public void run() {
        System.out.println("run方法开始");
        System.out.println(Thread.currentThread().isAlive());
        System.out.println("run方法运行状态：" + this.isAlive());
        System.out.println("run方法结束");
    }
}