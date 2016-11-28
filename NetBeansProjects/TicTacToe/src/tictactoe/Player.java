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
public class Player {
    private static Player p1=null;
    private static Player p2=null;
    private boolean num;
    private static boolean turn=false;
    private Player(boolean i){
        num=i;
    }
    public Player getP1(){
        if (p1==null){
            p1=new Player(false);
        }
        return p1;
    }
    public Player getP2(){
        if (p2==null){
            p2=new Player(true);
        }
        return p2;
    }
    public void move(int a){
        if (turn==this.num){
            if(Board.getInstance().getBoard()[a]==0){
                if(num){
                    Board.getInstance().updateBoard(a,1);//player one takes spot a
                }else{
                    Board.getInstance().updateBoard(a,2);//player two takes spot a
                }
            }else{
                //spot is filled
            }
        }else{
            //Not your turn
        }
        turn=!turn;
    }
}
