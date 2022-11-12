/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package Students;

/**
 *
 * @author THANG
 */
public class ObjectStu {
    private int Id;
    private String Name;
    private int Age;
    private String Address;
    private float Gpa;
    public ObjectStu(){
        
    }

    public ObjectStu(int Id, String Name, int Age, String Address, float Gpa) {
        super();
        this.Id = Id;
        this.Name = Name;
        this.Age = Age;
        this.Address = Address;
        this.Gpa = Gpa;
    }

    public int getId() {
        return Id;
    }

    public void setId(int Id) {
        this.Id = Id;
    }

    public String getName() {
        return Name;
    }

    public void setName(String Name) {
        this.Name = Name;
    }

    public int getAge() {
        return Age;
    }

    public void setAge(int Age) {
        this.Age = Age;
    }

    public String getAddress() {
        return Address;
    }

    public void setAddress(String Address) {
        this.Address = Address;
    }

    public float getGpa() {
        return Gpa;
    }

    public void setGpa(float Gpa) {
        this.Gpa = Gpa;
    }

    @Override
    public String toString() {
        return "ObjectStu{" + "Id=" + Id + ", Name=" + Name + ", Age=" + Age + ", Address=" + Address + ", Gpa=" + Gpa + '}';
    }
    
}
