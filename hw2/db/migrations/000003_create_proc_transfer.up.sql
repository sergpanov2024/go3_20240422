create or replace procedure procTransferTx (from_acc int8, to_acc int8, amount int8 )  
LANGUAGE plpgsql 
AS $$

begin

  INSERT INTO transfers        ( from_account_id,     to_account_id,    amount   ) 
                        VALUES ( from_acc,            to_acc,           amount   ) ;
                        
  INSERT INTO entries          ( account_id,   amount ) 
                        VALUES ( from_acc  ,  -amount ) ; 
                        
  INSERT INTO entries          ( account_id,   amount ) 
                        VALUES ( to_acc  ,  -amount ) ; 

  UPDATE accounts SET balance = balance - amount WHERE id = from_acc ;
  UPDATE accounts SET balance = balance + amount WHERE id = to_acc ;
  
  commit;

end ; 
$$
;