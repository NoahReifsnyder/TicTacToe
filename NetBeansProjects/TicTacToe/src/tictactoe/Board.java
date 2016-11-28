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
}
