package chap2;

public class Demo03 {
    public static void main(String[] args) {
        Demo03Service service1 = new Demo03Service();
        Demo03Service service2 = new Demo03Service();

        Thread t1 = new Demo03ThreadA(service1);
        Thread t2 = new Demo03ThreadB(service2);
        t1.start();
        t2.start();
    }
}

class Demo03Service {
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

class Demo03ThreadA extends Thread{
    private Demo03Service service;
    public Demo03ThreadA(Demo03Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.add("a");
    }
}

class Demo03ThreadB extends Thread{
    private Demo03Service service;
    public Demo03ThreadB(Demo03Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.add("b");
    }
}