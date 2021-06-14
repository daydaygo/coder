package chap2;

public class Demo11 {
    public static void main(String[] args) throws InterruptedException {
        Demo11Service service = new Demo11Service();
        Thread t1 = new Demo11ThreadA(service);
        t1.setName("A");
        t1.start();

        Thread t2 = new Demo11ThreadB(service);
        t2.setName("B");
        t2.start();

        Thread.sleep(10000);
        long start = Demo11Utils.start1 > Demo11Utils.start2 ? Demo11Utils.start2 : Demo11Utils.start1;
        long end = Demo11Utils.end1 > Demo11Utils.end2 ? Demo11Utils.end1 : Demo11Utils.end2;
        System.out.println("耗时：" + (end - start) /1000 + "秒");

    }
}

class Demo11Utils{
    static long start1;
    static long start2;
    static long end1;
    static long end2;
}

class Demo11Service{
    /*synchronized*/ public void foo(){
        try {
            System.out.println(Thread.currentThread().getName() + "开始任务");
            Thread.sleep(3000);
            synchronized (this) {
                System.out.println(Thread.currentThread().getName() + "处理计算结果");
            }
            System.out.println(Thread.currentThread().getName() + "结束任务");
        }catch(InterruptedException e){
            e.printStackTrace();
        }
    }
}

class Demo11ThreadA extends Thread{
    private Demo11Service service;
    public Demo11ThreadA(Demo11Service service){
        this.service = service;
    }

    @Override
    public void run() {
        Demo11Utils.start1 = System.currentTimeMillis();
        service.foo();
        Demo11Utils.end1 = System.currentTimeMillis();
    }
}

class Demo11ThreadB extends Thread{
    private Demo11Service service;
    public Demo11ThreadB(Demo11Service service){
        this.service = service;
    }

    @Override
    public void run() {
        Demo11Utils.start2 = System.currentTimeMillis();
        service.foo();
        Demo11Utils.end2 = System.currentTimeMillis();
    }
}