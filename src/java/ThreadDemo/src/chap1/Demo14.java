package chap1;

public class Demo14 {
    public static void main(String[] args) throws InterruptedException {
        Demo14User user = new Demo14User();
        Thread t = new Demo14Thread(user);
        t.start();
        Thread.sleep(500);
        t.stop();
        System.out.println("username=" + user.getUsername() + ", password=" + user.getPassword());
    }
}

class Demo14Thread extends Thread{
    private Demo14User user;
    public Demo14Thread(Demo14User user){
        this.user = user;
    }

    @Override
    public void run() {
        user.updateUsernameAndPassword("b", "bb");
    }
}

class Demo14User {
    private String username = "a";
    private String password = "aa";
    public void setUsername(String username){
        this.username = username;
    }
    public String getUsername(){
        return username;
    }
    public void setPassword(String passpword){
        this.password = passpword;
    }
    public String getPassword(){
        return password;
    }
    synchronized public void updateUsernameAndPassword(String username, String password){
        try {
            this.username = username;
            Thread.sleep(10000);
            this.password = password;
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}
