package tictactoe;

/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

/**
 *
 * @author Noah
 */
public class Board {
    private static Board b=null;
    private int[] board;
    private Board(){
        for(int i=0; i<9;i++){
            board[i]=0;
        }
    }
    public static Board getInstance(){
        if(b==null){
            b=new Board();
        }
        return b;
    }
    public int[] getBoard(){
        return board;
    }
    public void updateBoard(int a, int b){
        board[a]=b;
    }
    public boolean checkWin(Player p){
        if (!p.num){
            if (this.board[0]==this.board[1]&&this.board[1]==this.board[1]&&this.board[1]==1
                    || this.board[3]==this.board[4]&&this.board[4]==this.board[5]&&this.board[5]==1
                    || this.board[6]==this.board[7]&&this.board[7]==this.board[8]&&this.board[8]==1
                    || this.board[0]==this.board[3]&&this.board[3]==this.board[6]&&this.board[6]==1
                    || this.board[1]==this.board[4]&&this.board[4]==this.board[7]&&this.board[7]==1
                    || this.board[2]==this.board[5]&&this.board[5]==this.board[8]&&this.board[8]==1
                    || this.board[0]==this.board[4]&&this.board[4]==this.board[8]&&this.board[8]==1
                    || this.board[2]==this.board[4]&&this.board[4]==this.board[6]&&this.board[6]==1){
                
                return true;
            }
        }else{
            if (this.board[0]==this.board[1]&&this.board[1]==this.board[1]&&this.board[1]==2
                    || this.board[3]==this.board[4]&&this.board[4]==this.board[5]&&this.board[5]==2
                    || this.board[6]==this.board[7]&&this.board[7]==this.board[8]&&this.board[8]==2
                    || this.board[0]==this.board[3]&&this.board[3]==this.board[6]&&this.board[6]==2
                    || this.board[1]==this.board[4]&&this.board[4]==this.board[7]&&this.board[7]==2
                    || this.board[2]==this.board[5]&&this.board[5]==this.board[8]&&this.board[8]==2
                    || this.board[0]==this.board[4]&&this.board[4]==this.board[8]&&this.board[8]==2
                    || this.board[2]==this.board[4]&&this.board[4]==this.board[6]&&this.board[6]==2){
                
                return true;
            }
        }
        return false;
    }
}
