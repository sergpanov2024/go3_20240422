
create or replace function procTransferTx (from_acc int8, to_acc int8, amount int8 )  
returns table (Transfer int8,   FromAccount int8,  ToAccount int8,   FromEntry int8,   ToEntry int8)
LANGUAGE plpgsql 
AS $$
declare
  tr_id int8;
  ef_id int8;
  et_id int8;

begin

  INSERT INTO transfers        ( from_account_id,     to_account_id,    amount   ) 
                        VALUES ( from_acc,            to_acc,           amount   ) 
    returning id into tr_id;
                        
  INSERT INTO entries          ( account_id,   amount ) 
                        VALUES ( from_acc  ,  -amount ) 
    returning id into ef_id;
                        
  INSERT INTO entries          ( account_id,   amount ) 
                        VALUES ( to_acc  ,  -amount )                                 
    returning id into et_id;

  UPDATE accounts SET balance = balance - amount WHERE id = from_acc ;
  UPDATE accounts SET balance = balance + amount WHERE id = to_acc ;
  
  --commit;
   
  return query 
  select tr_id as Transfer, from_acc as FromAccount, to_acc as ToAccount, ef_id as FromEntry, et_id as ToEntry;

end ; 
$$
;


