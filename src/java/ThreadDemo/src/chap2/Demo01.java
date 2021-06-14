package chap2;

public class Demo01 {
    public static void main(String[] args) {
        Demo01Service service = new Demo01Service();
        Thread t1 = new Demo01ThreadA(service);
        t1.start();
        Thread t2 = new Demo01ThreadB(service);
        t2.start();
    }
}

class Demo01Service{
    public void add(String username){
        int num = 0;
        if ("a".equals(username)){
            num = 100;
            System.out.println("a set over");
            // 等待另外一个线程修改num的值
            try {
                Thread.sleep(2000);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }else {
            num = 200;
            System.out.println("b set over");
        }
        System.out.println("username=" + username + ", num=" + num);
    }
}

class Demo01ThreadA extends Thread{
    private Demo01Service service;
    public Demo01ThreadA(Demo01Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.add("a");
    }
}
class Demo01ThreadB extends Thread{
    private Demo01Service service;
    public Demo01ThreadB(Demo01Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.add("b");
    }
}