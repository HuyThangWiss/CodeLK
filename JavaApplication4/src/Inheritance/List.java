/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package Inheritance;

import java.util.ArrayList;
import java.util.Scanner;

/**
 *
 * @author THANG
 */
public class List {
    public static void main(String[] args) {
        ArrayList<Float> list=new ArrayList<Float>();
        int n;
        Scanner sc= new Scanner(System.in);
        System.out.println("Nhap so luong phan tu");
        n= Integer.parseInt(sc.nextLine());
        float k;
        for(int i=0;i<n;i++){
            System.out.printf("Nhap pt thu : ",(i+1));
            k=Float.parseFloat(sc.nextLine());
            list.add(k);
        }
        
        System.out.println(list);
    }
}
