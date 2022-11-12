/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package Inheritance;

import java.util.HashMap;
import java.util.Iterator;

/**
 *
 * @author THANG
 */
public class map {
    public static void main(String[] args) {
      //  HashMap<String,int> capitalCituties =new HashMap<String,int>();
//        HashMap<String, Integer> capitalCities = new HashMap<String, Integer>();
//        capitalCities.put("index 1", 1);
//        capitalCities.put("index 2", 2);
//        capitalCities.put("index 3", 3);
//        capitalCities.put("index 4", 4);
//        System.out.println(capitalCities.containsKey("index 1"));
//        System.out.println("Apter delete");
//        int retured_value = (int)capitalCities.remove("index 1");
//        System.out.println(capitalCities);
        HashMap<Langkey,String> langMap=new HashMap<Langkey,String>();
        langMap.put(new Langkey(1,"En"),"English");
        langMap.put(new Langkey(2, "VI"), "Vietnamese");
        langMap.put(new Langkey(3, "ES"), "Spainish");
        langMap.put(new Langkey(4, "JP"), "Jaspanese");
        
       System.out.println("Size map"+langMap.size());
      System.out.print("Cac phan tu cua langMap: ");
        Iterator<Langkey> itr = langMap.keySet().iterator();
        while (itr.hasNext()) {
            System.out.print(langMap.get(itr.next()) + " ");
        }
        System.out.println("\nSize cua langMap la: " + langMap.size());
   }
}
class Langkey{
    int index;
    String name;

    public Langkey(int index, String name) {
        this.index = index;
        this.name = name;
    }
     @Override
    public int hashCode(){
        return 5;
    }  
    @Override
    public boolean equals(Object obj){
        return true;
    }
}






