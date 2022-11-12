/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Main.java to edit this template
 */
package javaapplication4;

/**
 *
 * @author THANG
 */
public class JavaApplication4 {

   
    public static void main(String[] args) {
        
       Complex c1=new Complex(10,15);
        System.err.println(c1);
    }
    
}

class Complex{
    private double re,im;
    public Complex(){
        
    }

    public double getRe() {
        return re;
    }

    public void setRe(double re) {
        this.re = re;
    }

    public double getIm() {
        return im;
    }

    public void setIm(double im) {
        this.im = im;
    }

    public Complex(double re, double im) {
        this.re = re;
        this.im = im;
    }

    @Override
    public String toString() {
        return "Complex{" + "re=" + re + ", im=" + im + '}';
    }
    
}