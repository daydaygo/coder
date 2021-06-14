package chap2;

public class Demo02 {
    public static void main(String[] args) {
        Demo02Service service = new Demo02Service();
        Thread t1 = new Demo02ThreadA(service);
        Thread t2 = new Demo02ThreadB(service);
        t1.start();
        t2.start();
    }
}

class Demo02Service {
    private int num;
    synchronized public void add(String username){
        if ("a".equals(username)){
            num = 100;
            System.out.println("a is set");
            try {
                Thread.sleep(2000);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }else{
            num = 200;
            System.out.println("b is set");
        }
        System.out.println("username=" + username + ", num=" + num);
    }
}

class Demo02ThreadA extends Thread{
    private Demo02Service service;
    public Demo02ThreadA(Demo02Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.add("a");
    }
}

class Demo02ThreadB extends Thread{
    private Demo02Service service;
    public Demo02ThreadB(Demo02Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.add("b");
    }
}