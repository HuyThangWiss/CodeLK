/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package Students;

import java.util.Comparator;

/**
 *
 * @author THANG
 */

public class SortStudentByGPA implements Comparator<ObjectStu> {
    @Override
    public int compare(ObjectStu student1, ObjectStu student2) {
        if (student1.getGpa() > student2.getGpa()) {
            return 1;
        }
        return -1;
    }
}