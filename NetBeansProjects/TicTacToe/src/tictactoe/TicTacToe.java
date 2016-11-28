/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package tictactoe;

/**
 *
 * @author Noah
 */
public class TicTacToe {

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        Player p1=Player.getP1();
        Player p2=Player.getP2();
        Player currPlayer=p2;
        new BoardGUI();
        while (!Board.getInstance().checkWin(currPlayer)){
            if (currPlayer==p1){
                currPlayer=p2;
            }else{
                currPlayer=p1;
            }
            
            
            
        }
    }
    
}
