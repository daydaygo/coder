package chap1;

public class Demo22 {
    public static void main(String[] args) throws InterruptedException {
        Demo22User user = new Demo22User();
        Thread t1 = new Thread(){
            @Override
            public void run() {
                user.updateUsernameAndPassword("b", "bb");
            }
        };
        t1.setName("A");
        t1.start();

        Thread.sleep(10);

        Thread t2 = new Thread(){
            @Override
            public void run() {
                user.printUseruserAndPassword();
            }
        };
        t2.start();
    }
}

class Demo22User {
    private String username = "a";
    private String password = "aa";

    public void updateUsernameAndPassword(String username, String password){
        this.username = username;
        if ("A".equals(Thread.currentThread().getName())){
            System.out.println("停止A线程");
            Thread.currentThread().suspend();
        }
        this.password = password;
    }

    public void printUseruserAndPassword(){
        System.out.println("username=" + username + ", password=" +password);
    }
}