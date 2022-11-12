/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package Students;


import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.Scanner;

/**
 *
 * @author THANG
 */
public class StudentManager{
   public static  Scanner sc=new Scanner(System.in);
  // private List<ObjectStu> ListStu;
   private ArrayList<ObjectStu> ListStu=new ArrayList<>();
    public  StudentManager(){
        
    }

    public void AddStudents()
    {
        int n;
        System.out.println("Nhap so luong sinh vien : ");
        n=Integer.parseInt(sc.nextLine());
        for(int i=0;i<n;i++){          
                ObjectStu stu=new ObjectStu();              
                System.out.println("Nhap sv thu : "+(i+1));              
                System.out.println("Nhap id :");
                stu.setId(Integer.parseInt(sc.nextLine()));
                System.out.println("Nhap name :");
                stu.setName(sc.nextLine());
                System.out.println("Nhap Age :");
                stu.setAge(Integer.parseInt(sc.nextLine()));
                System.out.println("Nhap Address :");
                stu.setAddress(sc.nextLine());
                System.out.println("Nhap Gpa :");
                stu.setGpa(Float.parseFloat(sc.nextLine()));
                sc.nextLine();           
            ListStu.add(stu);         
        }       
    }
    public void Input(){
       int n;
        System.out.println("Nhap so luong sinh vien : ");
        n=Integer.parseInt(sc.nextLine());
        for(int i=0;i<n;i++){      
                System.out.println("Nhap sv thu : "+(i+1));              
                System.out.println("Nhap id :");
                int id = Integer.parseInt(sc.nextLine());
                System.out.println("Nhap name :");
                String name =sc.nextLine();
                System.out.println("Nhap Age :");
                int age = Integer.parseInt(sc.nextLine());
                System.out.println("Nhap Address :");
                String address = sc.nextLine();
                System.out.println("Nhap Gpa :");
                float Gpa =  Float.parseFloat(sc.nextLine());
                ObjectStu stu=new ObjectStu(id, name, age, address, Gpa);
                ListStu.add(stu);
                sc.nextLine();
        }    
    }
    public void Output()
    {
        System.out.println("Infor : "+(ListStu.toString()));
    }
    public void Show(){
        for(int i=0;i<ListStu.size();i++){
            System.out.println("Thong tin "+(ListStu.get(i).toString()));
            
        }
    }
    public void sortStudentByGPA() {
        Collections.sort(ListStu, new SortStudentByGPA());
    }
}

