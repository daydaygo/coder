package chap2;

public class Demo07 {
    public static void main(String[] args) {
        Demo07Service service = new Demo07Service();
        Thread t1 = new Demo07Thread(service);
        t1.setName("A");
        t1.start();
        Thread t2 = new Demo07Thread(service);
        t2.start();
    }
}

class Demo07Service {
    synchronized public void foo(){
        if ("A".equals(Thread.currentThread().getName())){
            System.out.println("线程A开始于" + System.currentTimeMillis());
            while(true){
                if(("" + Math.random()).substring(0, 8).equals("0.123456")){
                    System.out.println("线程A结束于" + System.currentTimeMillis());
                    Integer.parseInt("A");
                }
            }
        }else{
            System.out.println("线程B开始于" + System.currentTimeMillis());
        }
    }
}

class Demo07Thread extends Thread{
    private Demo07Service service;
    public Demo07Thread(Demo07Service service){
        this.service = service;
    }
    @Override
    public void run() {
        service.foo();
    }
}
