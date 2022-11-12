/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package Students;

/**
 *
 * @author THANG
 */
public class main {
    public static void main(String[] args) {
        StudentManager stu=new StudentManager();
     //   stu.AddStudents();
    //    stu.Show();
       
        stu.Input();
        System.out.println("*******************");
        stu.Output();
        System.out.println("*******************");
        stu.sortStudentByGPA();
        stu.Output();
    }
}
