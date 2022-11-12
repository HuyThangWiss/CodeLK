/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package Inheritance;

class Grandpa{
    public void show(){
        System.out.println("Inside show() method of Grandpa class");
    }
}
class Dad extends Grandpa{
    @Override public void show(){
         System.out.println(
            "Inside show() method of Dad class");
    }
}
class Me extends Dad{
    @Override public void show(){
        System.out.println(
            "Inside show() method of Me class");
    }
    
}
public class Inheritance {
    public static void main(String[] args) {
      Grandpa grandpa = new Grandpa();
      Grandpa dad =new Dad();
      Grandpa me = new Me();
      
      grandpa.show();
      dad.show();
      me.show();
    }
     
}
